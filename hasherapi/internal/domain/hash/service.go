package hash

import "context"

func NewService(
	calculator Calculator,
	storage Storage,
) *Service {
	return &Service{
		calculator: calculator,
		storage:    storage,
	}
}

type Service struct {
	calculator Calculator
	storage    Storage
}

func (s *Service) CreateHashes(ctx context.Context, inputs []Input) ([]IdentifiedSHA3Hash, error) {
	sha3Hashes, err := s.calculator.Calculate(ctx, inputs)
	if err != nil {
		return nil, err
	}

	identifiedSHA3Hashes, err := s.storage.Save(ctx, sha3Hashes)
	if err != nil {
		return nil, err
	}

	return identifiedSHA3Hashes, nil
}

func (s *Service) FindHashes(ctx context.Context, ids []ID) ([]IdentifiedSHA3Hash, error) {
	identifiedSHA3Hashes, err := s.storage.Get(ctx, ids)
	if err != nil {
		return nil, err
	}

	return identifiedSHA3Hashes, nil
}
