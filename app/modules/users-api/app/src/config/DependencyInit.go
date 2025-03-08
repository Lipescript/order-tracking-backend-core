package config

import (
   "users-api/app/src/controller"
   "users-api/app/src/repository"
   "users-api/app/src/service"
)

type Initialization struct {
   userRepo repository.UserRepository
   userSvc  service.UserService
   UserCtrl controller.UserController
}

func NewInitialization(
   userRepo repository.UserRepository,
   userService service.UserService,
   userCtrl controller.UserController *Initialization {
   return &Initialization{
      userRepo: userRepo,
      userSvc:  userService,
      UserCtrl: userCtrl,
   }
}