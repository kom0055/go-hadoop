package security

import (
	"github.com/kom0055/go-hadoop/proto/common"
	"log"
	"os/user"
	"sync"
)

/** a (very) basic UserGroupInformation implementation for storing user data/tokens,
  This implementation is currently *not* thread-safe
*/

type UserGroupInformation struct {
	rwMutex    sync.RWMutex
	userInfo   *common.UserInformationProto
	userTokens map[string]*common.TokenProto
}

var once sync.Once
var currentUserGroupInformation *UserGroupInformation
var maxTokens = 16

func CreateCurrentUserInfoProto() (*common.UserInformationProto, error) {
	// Figure the current user-name
	var username string
	if currentUser, err := user.Current(); err != nil {
		log.Fatal("user.Current", err)
		return nil, err
	} else {
		username = currentUser.Username
	}

	return &common.UserInformationProto{EffectiveUser: nil, RealUser: &username}, nil
}

func Allocate(userInfo *common.UserInformationProto, userTokens map[string]*common.TokenProto) *UserGroupInformation {
	ugi := new(UserGroupInformation)

	if userInfo != nil {
		ugi.userInfo = userInfo
	} else {
		currentUserInfo, _ := CreateCurrentUserInfoProto()
		ugi.userInfo = currentUserInfo
	}

	if userTokens != nil {
		ugi.userTokens = userTokens
	} else {
		ugi.userTokens = make(map[string]*common.TokenProto) //empty, with room for maxTokens tokens.
	}

	return ugi
}

func initializeCurrentUser() {
	once.Do(func() {
		currentUserGroupInformation = Allocate(nil, nil)
	})
}

func (ugi *UserGroupInformation) GetUserInformation() *common.UserInformationProto {
	return ugi.userInfo
}

func (ugi *UserGroupInformation) GetUserTokens() map[string]*common.TokenProto {
	return ugi.userTokens
}

func (ugi *UserGroupInformation) AddUserTokenWithAlias(alias string, token *common.TokenProto) {
	if token == nil {
		log.Fatal("supplied token is nil!")
		return
	}

	if length := len(ugi.userTokens); length < maxTokens {
		ugi.userTokens[alias] = token
	} else {
		log.Fatal("user already has maxTokens:", maxTokens)
	}
}

func (ugi *UserGroupInformation) AddUserToken(token *common.TokenProto) {
	if token == nil {
		log.Fatal("supplied token is nil!")
		return
	}

	ugi.AddUserTokenWithAlias(token.GetService(), token)
}

func GetCurrentUser() *UserGroupInformation {
	initializeCurrentUser()

	return currentUserGroupInformation
}
