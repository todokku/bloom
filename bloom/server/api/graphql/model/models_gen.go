// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AcceptGroupInvitationInput struct {
	// group id
	ID string `json:"id"`
}

type AddPaymentMethodInput struct {
	StripeID string `json:"stripeId"`
	// if groupId is null, add to current user
	GroupID *string `json:"groupId"`
}

type BillingPlanConnection struct {
	Edges      []*BillingPlanEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount Int64              `json:"totalCount"`
}

type BillingPlanEdge struct {
	Cursor string       `json:"cursor"`
	Node   *BillingPlan `json:"node"`
}

type BillingPlanInput struct {
	ID      *string        `json:"id"`
	Name    string         `json:"name"`
	Product BillingProduct `json:"product"`
	// the strip id of the stripe plan. starting with 'plan_'
	StripeID string `json:"stripeId"`
	// HTML description
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	Storage     Int64  `json:"storage"`
}

type BillingSubscription struct {
	UpdatedAt            time.Time    `json:"updatedAt"`
	UsedStorage          Int64        `json:"usedStorage"`
	StripeCustomerID     *string      `json:"stripeCustomerId"`
	StripeSubscriptionID *string      `json:"stripeSubscriptionId"`
	Plan                 *BillingPlan `json:"plan"`
}

type BloomMetadata struct {
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
}

type CancelGroupInvitationInput struct {
	// group id
	ID string `json:"id"`
}

// set payment method with `id` as the default one
type ChangeDefaultPaymentMethodInput struct {
	ID string `json:"id"`
}

type CompleteRegistrationInput struct {
	// pending user id
	ID                  string              `json:"id"`
	Username            string              `json:"username"`
	AuthKey             Bytes               `json:"authKey"`
	Device              *SessionDeviceInput `json:"device"`
	PublicKey           Bytes               `json:"publicKey"`
	EncryptedPrivateKey Bytes               `json:"encryptedPrivateKey"`
}

type CreateGroupInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// users to invite, by username
	UsersToInvite []string `json:"usersToInvite"`
}

type DeclineGroupInvitationInput struct {
	// group id
	ID string `json:"id"`
}

type DeleteBillingPlanInput struct {
	ID string `json:"id"`
}

type DeleteGroupInput struct {
	ID string `json:"id"`
}

type GroupConnection struct {
	Edges      []*GroupEdge `json:"edges"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount Int64        `json:"totalCount"`
}

type GroupEdge struct {
	Cursor string `json:"cursor"`
	Node   *Group `json:"node"`
}

type GroupInput struct {
	// group id
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GroupInvitation struct {
	ID      string `json:"id"`
	Group   *Group `json:"group"`
	Inviter *User  `json:"inviter"`
	Invitee *User  `json:"invitee"`
}

type GroupInvitationConnection struct {
	Edges      []*GroupInvitationEdge `json:"edges"`
	PageInfo   *PageInfo              `json:"pageInfo"`
	TotalCount Int64                  `json:"totalCount"`
}

type GroupInvitationEdge struct {
	Cursor string           `json:"cursor"`
	Node   *GroupInvitation `json:"node"`
}

type GroupMemberConnection struct {
	Edges      []*GroupMemberEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount Int64              `json:"totalCount"`
}

type GroupMemberEdge struct {
	Cursor   string           `json:"cursor"`
	Node     *User            `json:"node"`
	Role     *GroupMemberRole `json:"role"`
	JoinedAt *time.Time       `json:"joinedAt"`
}

type InviteUsersInGroupInput struct {
	// group id
	ID string `json:"id"`
	// users to invite, by username
	Users []string `json:"users"`
}

type Invoice struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Amount          Int64     `json:"amount"`
	StripeID        string    `json:"stripeId"`
	StripeHostedURL string    `json:"stripeHostedUrl"`
	StripePdfURL    string    `json:"stripePdfUrl"`
	Paid            bool      `json:"paid"`
}

type InvoiceConnection struct {
	Edges      []*InvoiceEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount Int64          `json:"totalCount"`
}

type InvoiceEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Invoice `json:"node"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

type PaymentMethod struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"createdAt"`
	CardLast4           string    `json:"cardLast4"`
	CardExpirationMonth int       `json:"cardExpirationMonth"`
	CardExpirationYear  int       `json:"cardExpirationYear"`
	IsDefault           bool      `json:"isDefault"`
}

type PaymentMethodConnection struct {
	Edges      []*PaymentMethodEdge `json:"edges"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount Int64                `json:"totalCount"`
}

type PaymentMethodEdge struct {
	Cursor string         `json:"cursor"`
	Node   *PaymentMethod `json:"node"`
}

type QuitGroupInput struct {
	// group id
	ID string `json:"id"`
}

type RegisterInput struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type RegistrationStarted struct {
	ID string `json:"id"`
}

type RemoveGroupMembersInput struct {
	// group id
	ID string `json:"id"`
	// members to remvove, by username
	Members []string `json:"members"`
}

// remove payment method with `id`
type RemovePaymentMethodInput struct {
	ID string `json:"id"`
}

type RevokeSessionInput struct {
	ID string `json:"id"`
}

type SendNewRegistrationCodeInput struct {
	ID string `json:"id"`
}

type Session struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Token     *string        `json:"token"`
	Device    *SessionDevice `json:"device"`
}

type SessionConnection struct {
	Edges      []*SessionEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount Int64          `json:"totalCount"`
}

type SessionDevice struct {
	Os   SessionDeviceOs   `json:"os"`
	Type SessionDeviceType `json:"type"`
}

type SessionDeviceInput struct {
	Os   SessionDeviceOs   `json:"os"`
	Type SessionDeviceType `json:"type"`
}

type SessionEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Session `json:"node"`
}

type SignInInput struct {
	Username string              `json:"username"`
	AuthKey  Bytes               `json:"authKey"`
	Device   *SessionDeviceInput `json:"device"`
}

type SignedIn struct {
	Session *Session `json:"session"`
	Me      *User    `json:"me"`
}

// if groupId and userId are null (reserved for admins), add to current user
type UpdateBillingSubscriptionInput struct {
	PlanID  string  `json:"planId"`
	UserID  *string `json:"userId"`
	GroupID *string `json:"groupId"`
}

type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount Int64       `json:"totalCount"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type UserProfileInput struct {
	// id is reserved for admins
	ID          *string `json:"id"`
	DisplayName *string `json:"displayName"`
	Bio         *string `json:"bio"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
}

type VerifyRegistrationInput struct {
	// pending user id
	ID   string `json:"id"`
	Code string `json:"code"`
}

type BillingProduct string

const (
	BillingProductFree  BillingProduct = "FREE"
	BillingProductLite  BillingProduct = "LITE"
	BillingProductPro   BillingProduct = "PRO"
	BillingProductUltra BillingProduct = "ULTRA"
)

var AllBillingProduct = []BillingProduct{
	BillingProductFree,
	BillingProductLite,
	BillingProductPro,
	BillingProductUltra,
}

func (e BillingProduct) IsValid() bool {
	switch e {
	case BillingProductFree, BillingProductLite, BillingProductPro, BillingProductUltra:
		return true
	}
	return false
}

func (e BillingProduct) String() string {
	return string(e)
}

func (e *BillingProduct) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BillingProduct(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BillingProduct", str)
	}
	return nil
}

func (e BillingProduct) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GroupMemberRole string

const (
	GroupMemberRoleAdmin  GroupMemberRole = "ADMIN"
	GroupMemberRoleMember GroupMemberRole = "MEMBER"
)

var AllGroupMemberRole = []GroupMemberRole{
	GroupMemberRoleAdmin,
	GroupMemberRoleMember,
}

func (e GroupMemberRole) IsValid() bool {
	switch e {
	case GroupMemberRoleAdmin, GroupMemberRoleMember:
		return true
	}
	return false
}

func (e GroupMemberRole) String() string {
	return string(e)
}

func (e *GroupMemberRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GroupMemberRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GroupMemberRole", str)
	}
	return nil
}

func (e GroupMemberRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionDeviceOs string

const (
	SessionDeviceOsLinux   SessionDeviceOs = "LINUX"
	SessionDeviceOsMacos   SessionDeviceOs = "MACOS"
	SessionDeviceOsWindows SessionDeviceOs = "WINDOWS"
	SessionDeviceOsAndroid SessionDeviceOs = "ANDROID"
	SessionDeviceOsIos     SessionDeviceOs = "IOS"
	SessionDeviceOsOther   SessionDeviceOs = "OTHER"
)

var AllSessionDeviceOs = []SessionDeviceOs{
	SessionDeviceOsLinux,
	SessionDeviceOsMacos,
	SessionDeviceOsWindows,
	SessionDeviceOsAndroid,
	SessionDeviceOsIos,
	SessionDeviceOsOther,
}

func (e SessionDeviceOs) IsValid() bool {
	switch e {
	case SessionDeviceOsLinux, SessionDeviceOsMacos, SessionDeviceOsWindows, SessionDeviceOsAndroid, SessionDeviceOsIos, SessionDeviceOsOther:
		return true
	}
	return false
}

func (e SessionDeviceOs) String() string {
	return string(e)
}

func (e *SessionDeviceOs) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionDeviceOs(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionDeviceOS", str)
	}
	return nil
}

func (e SessionDeviceOs) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionDeviceType string

const (
	SessionDeviceTypeTv       SessionDeviceType = "TV"
	SessionDeviceTypeConsole  SessionDeviceType = "CONSOLE"
	SessionDeviceTypeMobile   SessionDeviceType = "MOBILE"
	SessionDeviceTypeTablet   SessionDeviceType = "TABLET"
	SessionDeviceTypeWatch    SessionDeviceType = "WATCH"
	SessionDeviceTypeComputer SessionDeviceType = "COMPUTER"
	SessionDeviceTypeCar      SessionDeviceType = "CAR"
	SessionDeviceTypeOther    SessionDeviceType = "OTHER"
)

var AllSessionDeviceType = []SessionDeviceType{
	SessionDeviceTypeTv,
	SessionDeviceTypeConsole,
	SessionDeviceTypeMobile,
	SessionDeviceTypeTablet,
	SessionDeviceTypeWatch,
	SessionDeviceTypeComputer,
	SessionDeviceTypeCar,
	SessionDeviceTypeOther,
}

func (e SessionDeviceType) IsValid() bool {
	switch e {
	case SessionDeviceTypeTv, SessionDeviceTypeConsole, SessionDeviceTypeMobile, SessionDeviceTypeTablet, SessionDeviceTypeWatch, SessionDeviceTypeComputer, SessionDeviceTypeCar, SessionDeviceTypeOther:
		return true
	}
	return false
}

func (e SessionDeviceType) String() string {
	return string(e)
}

func (e *SessionDeviceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionDeviceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionDeviceType", str)
	}
	return nil
}

func (e SessionDeviceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
