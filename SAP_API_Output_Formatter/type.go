package sap_api_output_formatter

type Lead struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	LeadCode      string `json:"lead_code"`
	Deleted       bool   `json:"deleted"`
}

type LeadCollection struct {
	ObjectID                                                     string `json:"ObjectID"`
	ID                                                           string `json:"ID"`
	Name                                                         string `json:"Name"`
	NameLanguageCode                                             string `json:"NameLanguageCode"`
	AccountPartyID                                               string `json:"AccountPartyID"`
	AccountPartyUUID                                             string `json:"AccountPartyUUID"`
	AccountPartyName                                             string `json:"AccountPartyName"`
	ContactID                                                    string `json:"ContactID"`
	ContactUUID                                                  string `json:"ContactUUID"`
	ContactName                                                  string `json:"ContactName"`
	Company                                                      string `json:"Company"`
	ContactFirstName                                             string `json:"ContactFirstName"`
	ContactMiddleName                                            string `json:"ContactMiddleName"`
	ContactLastName                                              string `json:"ContactLastName"`
	IndividualCustomerGivenName                                  string `json:"IndividualCustomerGivenName"`
	IndividualCustomerMiddleName                                 string `json:"IndividualCustomerMiddleName"`
	IndividualCustomerFamilyName                                 string `json:"IndividualCustomerFamilyName"`
	QualificationLevelCode                                       string `json:"QualificationLevelCode"`
	UserStatusCode                                               string `json:"UserStatusCode"`
	ResultReasonCode                                             string `json:"ResultReasonCode"`
	ApprovalStatusCode                                           string `json:"ApprovalStatusCode"`
	ConsistencyStatusCode                                        string `json:"ConsistencyStatusCode"`
	OriginTypeCode                                               string `json:"OriginTypeCode"`
	PriorityCode                                                 string `json:"PriorityCode"`
	StartDate                                                    string `json:"StartDate"`
	EndDate                                                      string `json:"EndDate"`
	CampaignID                                                   string `json:"CampaignID"`
	GroupCode                                                    string `json:"GroupCode"`
	OwnerPartyID                                                 string `json:"OwnerPartyID"`
	OwnerPartyUUID                                               string `json:"OwnerPartyUUID"`
	EmployeeResponsibleUUID                                      string `json:"EmployeeResponsibleUUID"`
	OwnerPartyName                                               string `json:"OwnerPartyName"`
	OwnerPartyIDSales                                            string `json:"OwnerPartyIDSales"`
	OwnerUUIDSales                                               string `json:"OwnerUUIDSales"`
	OwnerSalesName                                               string `json:"OwnerSalesName"`
	MarketingUnitPartyID                                         string `json:"MarketingUnitPartyID"`
	MarketingUnitPartyUUID                                       string `json:"MarketingUnitPartyUUID"`
	MarketingName                                                string `json:"MarketingName"`
	SalesUnitPartyID                                             string `json:"SalesUnitPartyID"`
	SalesUnitPartyUUID                                           string `json:"SalesUnitPartyUUID"`
	SalesUnitName                                                string `json:"SalesUnitName"`
	SalesOrganisationID                                          string `json:"SalesOrganisationID"`
	SalesOrganisationUUID                                        string `json:"SalesOrganisationUUID"`
	DistributionChannelCode                                      string `json:"DistributionChannelCode"`
	DivisionCode                                                 string `json:"DivisionCode"`
	SalesOfficeID                                                string `json:"SalesOfficeID"`
	SalesOfficeUUID                                              string `json:"SalesOfficeUUID"`
	SalesGroupID                                                 string `json:"SalesGroupID"`
	SalesGroupUUID                                               string `json:"SalesGroupUUID"`
	SurveyTotalScoreValue                                        int    `json:"SurveyTotalScoreValue"`
	CreationDateTime                                             string `json:"CreationDateTime"`
	LastChangeDateTime                                           string `json:"LastChangeDateTime"`
	CreationIdentityUUID                                         string `json:"CreationIdentityUUID"`
	LastChangeIdentityUUID                                       string `json:"LastChangeIdentityUUID"`
	Note                                                         string `json:"Note"`
	SalesTerritoryID                                             string `json:"SalesTerritoryID"`
	SalesTerritoryUUID                                           string `json:"SalesTerritoryUUID"`
	SalesTerritoryName                                           string `json:"SalesTerritoryName"`
	ExpectedRevenueAmount                                        string `json:"ExpectedRevenueAmount"`
	ExpectedRevenueCurrencyCode                                  string `json:"ExpectedRevenueCurrencyCode"`
	Score                                                        int    `json:"Score"`
	CompanySecondName                                            string `json:"CompanySecondName"`
	CompanyThirdName                                             string `json:"CompanyThirdName"`
	CompanyFourthName                                            string `json:"CompanyFourthName"`
	AccountPostalAddressElementsHouseID                          string `json:"AccountPostalAddressElementsHouseID"`
	AccountPostalAddressElementsStreetPrefix                     string `json:"AccountPostalAddressElementsStreetPrefix"`
	AccountPostalAddressElementsAdditionalStreetPrefixName       string `json:"AccountPostalAddressElementsAdditionalStreetPrefixName"`
	AccountPostalAddressElementsStreetName                       string `json:"AccountPostalAddressElementsStreetName"`
	AccountPostalAddressElementsStreetSufix                      string `json:"AccountPostalAddressElementsStreetSufix"`
	AccountPostalAddressElementsAdditionalStreetSuffixName       string `json:"AccountPostalAddressElementsAdditionalStreetSuffixName"`
	AccountCity                                                  string `json:"AccountCity"`
	AccountCountry                                               string `json:"AccountCountry"`
	AccountState                                                 string `json:"AccountState"`
	AccountPostalAddressElementsPOBoxID                          string `json:"AccountPostalAddressElementsPOBoxID"`
	AccountPostalAddressElementsStreetPostalCode                 string `json:"AccountPostalAddressElementsStreetPostalCode"`
	AccountPostalAddressElementsCountyName                       string `json:"AccountPostalAddressElementsCountyName"`
	AccountPhone                                                 string `json:"AccountPhone"`
	AccountFax                                                   string `json:"AccountFax"`
	AccountMobile                                                string `json:"AccountMobile"`
	AccountEMail                                                 string `json:"AccountEMail"`
	AccountWebsite                                               string `json:"AccountWebsite"`
	AccountLatitudeMeasure                                       string `json:"AccountLatitudeMeasure"`
	AccountLatitudeMeasureUnitCode                               string `json:"AccountLatitudeMeasureUnitCode"`
	AccountLongitudeMeasure                                      string `json:"AccountLongitudeMeasure"`
	AccountLongitudeMeasureUnitCode                              string `json:"AccountLongitudeMeasureUnitCode"`
	AccountLegalForm                                             string `json:"AccountLegalForm"`
	OrganisationAccountABCClassificationCode                     string `json:"OrganisationAccountABCClassificationCode"`
	OrganisationAccountIndustrialSectorCode                      string `json:"OrganisationAccountIndustrialSectorCode"`
	AccountDUNS                                                  string `json:"AccountDUNS"`
	OrganisationAccountContactAllowedCode                        string `json:"OrganisationAccountContactAllowedCode"`
	AccountCorrespondenceLanguageCode                            string `json:"AccountCorrespondenceLanguageCode"`
	AccountNote                                                  string `json:"AccountNote"`
	ContactFormOfAddressCode                                     string `json:"ContactFormOfAddressCode"`
	ContactFunctionalTitleName                                   string `json:"ContactFunctionalTitleName"`
	ContactAcademicTitleCode                                     string `json:"ContactAcademicTitleCode"`
	ContactAdditionalAcademicTitleCode                           string `json:"ContactAdditionalAcademicTitleCode"`
	ContactNickName                                              string `json:"ContactNickName"`
	ContactCorrespondenceLanguageCode                            string `json:"ContactCorrespondenceLanguageCode"`
	ContactGenderCode                                            string `json:"ContactGenderCode"`
	ContactMaritalStatusCode                                     string `json:"ContactMaritalStatusCode"`
	BusinessPartnerRelationshipBusinessPartnerFunctionTypeCode   string `json:"BusinessPartnerRelationshipBusinessPartnerFunctionTypeCode"`
	BusinessPartnerRelationshipBusinessPartnerFunctionalAreaCode string `json:"BusinessPartnerRelationshipBusinessPartnerFunctionalAreaCode"`
	ContactDepartmentName                                        string `json:"ContactDepartmentName"`
	BusinessPartnerRelationshipContactVIPReasonCode              string `json:"BusinessPartnerRelationshipContactVIPReasonCode"`
	ContactAllowedCode                                           string `json:"ContactAllowedCode"`
	BusinessPartnerRelationshipEngagementScoreNumberValue        string `json:"BusinessPartnerRelationshipEngagementScoreNumberValue"`
	ContactBuildingID                                            string `json:"ContactBuildingID"`
	ContactFloorID                                               string `json:"ContactFloorID"`
	ContactRoomID                                                string `json:"ContactRoomID"`
	ContactPhone                                                 string `json:"ContactPhone"`
	ContactFacsimileFormattedNumberDescription                   string `json:"ContactFacsimileFormattedNumberDescription"`
	ContactMobile                                                string `json:"ContactMobile"`
	ContactEMail                                                 string `json:"ContactEMail"`
	ContactEMailUsageDeniedIndicator                             string `json:"ContactEMailUsageDeniedIndicator"`
	ContactNote                                                  string `json:"ContactNote"`
	IndividualCustomerABCClassificationCode                      string `json:"IndividualCustomerABCClassificationCode"`
	IndividualCustomerGenderCode                                 string `json:"IndividualCustomerGenderCode"`
	IndividualCustomerMaritalStatusCode                          string `json:"IndividualCustomerMaritalStatusCode"`
	IndividualCustomerEMail                                      string `json:"IndividualCustomerEMail"`
	IndividualCustomerPhone                                      string `json:"IndividualCustomerPhone"`
	IndividualCustomerFax                                        string `json:"IndividualCustomerFax"`
	IndividualCustomerMobile                                     string `json:"IndividualCustomerMobile"`
	IndividualCustomerAddressHouseID                             string `json:"IndividualCustomerAddressHouseID"`
	IndividualCustomerAddressStreetName                          string `json:"IndividualCustomerAddressStreetName"`
	IndividualCustomerAddressCity                                string `json:"IndividualCustomerAddressCity"`
	IndividualCustomerAddressCountry                             string `json:"IndividualCustomerAddressCountry"`
	IndividualCustomerAddressState                               string `json:"IndividualCustomerAddressState"`
	IndividualCustomerAddressPostalCode                          string `json:"IndividualCustomerAddressPostalCode"`
	IndividualCustomerAddressCountyName                          string `json:"IndividualCustomerAddressCountyName"`
	IndividualCustomerNationalityCountryCode                     string `json:"IndividualCustomerNationalityCountryCode"`
	IndividualCustomerBirthDate                                  string `json:"IndividualCustomerBirthDate"`
	IndividualCustomerFormOfAddressCode                          string `json:"IndividualCustomerFormOfAddressCode"`
	IndividualCustomerAcademicTitleCode                          string `json:"IndividualCustomerAcademicTitleCode"`
	IndividualCustomerOccupationCode                             string `json:"IndividualCustomerOccupationCode"`
	IndividualCustomerContactAllowedCode                         string `json:"IndividualCustomerContactAllowedCode"`
	IndividualCustomerNonVerbalCommunicationLanguageCode         string `json:"IndividualCustomerNonVerbalCommunicationLanguageCode"`
	IndividualCustomerInitialsName                               string `json:"IndividualCustomerInitialsName"`
	AccountPreferredCommunicationMediumTypeCode                  string `json:"AccountPreferredCommunicationMediumTypeCode"`
	IndividualCustomerNamePrefixCode                             string `json:"IndividualCustomerNamePrefixCode"`
	AccountLifeCycleStatusCode                                   string `json:"AccountLifeCycleStatusCode"`
	MainContactLifeCycleStatusCode                               string `json:"MainContactLifeCycleStatusCode"`
	LifeCycleStatusCode                                          string `json:"LifeCycleStatusCode"`
	ExternalID                                                   string `json:"ExternalID"`
	CreatedBy                                                    string `json:"CreatedBy"`
	LastChangedBy                                                string `json:"LastChangedBy"`
	EntityLastChangedOn                                          string `json:"EntityLastChangedOn"`
	ETag                                                         string `json:"ETag"`
	ToLeadItemCollection                                         string `json:"LeadItemCollection"`
}

type LeadItemCollection struct {
	ObjectID                   string `json:"ObjectID"`
	ParentObjectID             string `json:"ParentObjectID"`
	LeadID                     string `json:"LeadID"`
	ID                         string `json:"ID"`
	UUID                       string `json:"UUID"`
	Description                string `json:"Description"`
	LanguageCode               string `json:"languageCode"`
	LanguageCodeText           string `json:"languageCodeText"`
	ProductID                  string `json:"ProductID"`
	ProductUUID                string `json:"ProductUUID"`
	ProductName                string `json:"ProductName"`
	ProductCategoryInternalID  string `json:"ProductCategoryInternalID"`
	Quantity                   string `json:"Quantity"`
	UnitCode                   string `json:"unitCode"`
	UnitCodeText               string `json:"unitCodeText"`
	ProductCategoryDescription string `json:"ProductCategoryDescription"`
	EntityLastChangedOn        string `json:"EntityLastChangedOn"`
	ETag                       string `json:"ETag"`
}