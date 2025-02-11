/* Copyright (c) 2021, VRAI Labs and/or its affiliates. All rights reserved.
 *
 * This software is licensed under the Apache License, Version 2.0 (the
 * "License") as published by the Apache Software Foundation.
 *
 * You may not use this file except in compliance with the License. You may
 * obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package thirdpartyemailpassword

import (
	"errors"

	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Init(config *tpepmodels.TypeInput) supertokens.Recipe {
	return recipeInit(config)
}

func ThirdPartySignInUpWithContext(thirdPartyID string, thirdPartyUserID string, email tpepmodels.EmailStruct, userContext supertokens.UserContext) (tpepmodels.SignInUpResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return tpepmodels.SignInUpResponse{}, err
	}
	return (*instance.RecipeImpl.ThirdPartySignInUp)(thirdPartyID, thirdPartyUserID, email, userContext)
}

func GetUserByThirdPartyInfoWithContext(thirdPartyID string, thirdPartyUserID string, email tpmodels.EmailStruct, userContext supertokens.UserContext) (*tpepmodels.User, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.GetUserByThirdPartyInfo)(thirdPartyID, thirdPartyUserID, userContext)
}

func EmailPasswordSignUpWithContext(email, password string, userContext supertokens.UserContext) (tpepmodels.SignUpResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return tpepmodels.SignUpResponse{}, err
	}
	return (*instance.RecipeImpl.EmailPasswordSignUp)(email, password, userContext)
}

func EmailPasswordSignInWithContext(email, password string, userContext supertokens.UserContext) (tpepmodels.SignInResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return tpepmodels.SignInResponse{}, err
	}
	return (*instance.RecipeImpl.EmailPasswordSignIn)(email, password, userContext)
}

func GetUserByIdWithContext(userID string, userContext supertokens.UserContext) (*tpepmodels.User, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.GetUserByID)(userID, userContext)
}

func GetUsersByEmailWithContext(email string, userContext supertokens.UserContext) ([]tpepmodels.User, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.GetUsersByEmail)(email, userContext)
}

func CreateResetPasswordTokenWithContext(userID string, userContext supertokens.UserContext) (epmodels.CreateResetPasswordTokenResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return epmodels.CreateResetPasswordTokenResponse{}, err
	}
	return (*instance.RecipeImpl.CreateResetPasswordToken)(userID, userContext)
}

func ResetPasswordUsingTokenWithContext(token, newPassword string, userContext supertokens.UserContext) (epmodels.ResetPasswordUsingTokenResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return epmodels.ResetPasswordUsingTokenResponse{}, err
	}
	return (*instance.RecipeImpl.ResetPasswordUsingToken)(token, newPassword, userContext)
}

func UpdateEmailOrPasswordWithContext(userId string, email *string, password *string, userContext supertokens.UserContext) (epmodels.UpdateEmailOrPasswordResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return epmodels.UpdateEmailOrPasswordResponse{}, err
	}
	return (*instance.RecipeImpl.UpdateEmailOrPassword)(userId, email, password, userContext)
}

func CreateEmailVerificationTokenWithContext(userID string, userContext supertokens.UserContext) (evmodels.CreateEmailVerificationTokenResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return evmodels.CreateEmailVerificationTokenResponse{}, err
	}
	email, err := instance.getEmailForUserId(userID, userContext)
	if err != nil {
		return evmodels.CreateEmailVerificationTokenResponse{}, err
	}
	return (*instance.EmailVerificationRecipe.RecipeImpl.CreateEmailVerificationToken)(userID, email, userContext)
}

func VerifyEmailUsingTokenWithContext(token string, userContext supertokens.UserContext) (*tpepmodels.User, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	response, err := (*instance.EmailVerificationRecipe.RecipeImpl.VerifyEmailUsingToken)(token, userContext)
	if err != nil {
		return nil, err
	}
	if response.EmailVerificationInvalidTokenError != nil {
		return nil, errors.New("email verification token is invalid")
	}
	return (*instance.RecipeImpl.GetUserByID)(response.OK.User.ID, userContext)
}

func IsEmailVerifiedWithContext(userID string, userContext supertokens.UserContext) (bool, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return false, err
	}
	email, err := instance.getEmailForUserId(userID, userContext)
	if err != nil {
		return false, err
	}
	return (*instance.EmailVerificationRecipe.RecipeImpl.IsEmailVerified)(userID, email, userContext)
}

func RevokeEmailVerificationTokensWithContext(userID string, userContext supertokens.UserContext) (evmodels.RevokeEmailVerificationTokensResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return evmodels.RevokeEmailVerificationTokensResponse{}, err
	}
	email, err := instance.getEmailForUserId(userID, userContext)
	if err != nil {
		return evmodels.RevokeEmailVerificationTokensResponse{}, err
	}
	return (*instance.EmailVerificationRecipe.RecipeImpl.RevokeEmailVerificationTokens)(userID, email, userContext)
}

func UnverifyEmailWithContext(userID string, userContext supertokens.UserContext) (evmodels.UnverifyEmailResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return evmodels.UnverifyEmailResponse{}, err
	}
	email, err := instance.getEmailForUserId(userID, userContext)
	if err != nil {
		return evmodels.UnverifyEmailResponse{}, err
	}
	return (*instance.EmailVerificationRecipe.RecipeImpl.UnverifyEmail)(userID, email, userContext)
}

func ThirdPartySignInUp(thirdPartyID string, thirdPartyUserID string, email tpepmodels.EmailStruct) (tpepmodels.SignInUpResponse, error) {
	return ThirdPartySignInUpWithContext(thirdPartyID, thirdPartyUserID, email, &map[string]interface{}{})
}

func GetUserByThirdPartyInfo(thirdPartyID string, thirdPartyUserID string, email tpmodels.EmailStruct) (*tpepmodels.User, error) {
	return GetUserByThirdPartyInfoWithContext(thirdPartyID, thirdPartyUserID, email, &map[string]interface{}{})
}

func EmailPasswordSignUp(email, password string) (tpepmodels.SignUpResponse, error) {
	return EmailPasswordSignUpWithContext(email, password, &map[string]interface{}{})
}

func EmailPasswordSignIn(email, password string) (tpepmodels.SignInResponse, error) {
	return EmailPasswordSignInWithContext(email, password, &map[string]interface{}{})
}

func GetUserById(userID string) (*tpepmodels.User, error) {
	return GetUserByIdWithContext(userID, &map[string]interface{}{})
}

func GetUsersByEmail(email string) ([]tpepmodels.User, error) {
	return GetUsersByEmailWithContext(email, &map[string]interface{}{})
}

func CreateResetPasswordToken(userID string) (epmodels.CreateResetPasswordTokenResponse, error) {
	return CreateResetPasswordTokenWithContext(userID, &map[string]interface{}{})
}

func ResetPasswordUsingToken(token, newPassword string) (epmodels.ResetPasswordUsingTokenResponse, error) {
	return ResetPasswordUsingTokenWithContext(token, newPassword, &map[string]interface{}{})
}

func UpdateEmailOrPassword(userId string, email *string, password *string) (epmodels.UpdateEmailOrPasswordResponse, error) {
	return UpdateEmailOrPasswordWithContext(userId, email, password, &map[string]interface{}{})
}

func CreateEmailVerificationToken(userID string) (evmodels.CreateEmailVerificationTokenResponse, error) {
	return CreateEmailVerificationTokenWithContext(userID, &map[string]interface{}{})
}

func VerifyEmailUsingToken(token string) (*tpepmodels.User, error) {
	return VerifyEmailUsingTokenWithContext(token, &map[string]interface{}{})
}

func IsEmailVerified(userID string) (bool, error) {
	return IsEmailVerifiedWithContext(userID, &map[string]interface{}{})
}

func RevokeEmailVerificationTokens(userID string) (evmodels.RevokeEmailVerificationTokensResponse, error) {
	return RevokeEmailVerificationTokensWithContext(userID, &map[string]interface{}{})
}

func UnverifyEmail(userID string) (evmodels.UnverifyEmailResponse, error) {
	return UnverifyEmailWithContext(userID, &map[string]interface{}{})
}
