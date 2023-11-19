package common

func VerifyTenantID(tenantId string) bool {

	if tenantId == "" {
		return false
	}
	return true
}
