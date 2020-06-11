/**
 * Created by zc on 2020/6/10.
 */
package service

type Service struct {}

func New() Interface {
	return &Service{}
}

func (s *Service) User() UserService {
	return &userService{}
}

func (s *Service) Space() SpaceService {
	return &spaceService{}
}