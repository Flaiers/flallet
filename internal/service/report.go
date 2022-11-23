package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/flaiers/flallet/internal/dto"
	"github.com/flaiers/flallet/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/s3"
	"gorm.io/gorm"
)

type ReportService struct {
	repository         *repository.S3Repository
	transactionService *TransactionService
}

func NewReportService(db *gorm.DB, storage *s3.Storage) *ReportService {
	return &ReportService{
		repository:         repository.NewS3Repository(storage),
		transactionService: NewTransactionService(db),
	}
}

func (s *ReportService) GetCSV(reportFilter *dto.ReportFilter) (string, error) {
	report, err := s.transactionService.Report(
		reportFilter.Month, reportFilter.Year,
	)
	if err != nil {
		return "", err
	}

	buffer := &bytes.Buffer{}
	var records [][]string
	for _, row := range report {
		records = append(records, row.ToString())
	}

	w := csv.NewWriter(buffer)
	if err := w.WriteAll(records); err != nil {
		return "", err
	}
	if buffer.Len() == 0 {
		return "", fiber.NewError(fiber.StatusNotFound, "report not found")
	}

	key := fmt.Sprintf("%s.csv", time.Now().Format("2006-01-02-15-04-05"))
	url, err := s.repository.Upload(key, buffer.Bytes())
	if err != nil {
		return "", err
	}

	return url, nil
}
