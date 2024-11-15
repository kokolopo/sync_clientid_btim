package sql

var GETUSERBLINK = `
	SELECT
		kyc.ClientID, kyc.Email, kyc.IdentityNo
	FROM
		client_tbl_kyc kyc
	LIMIT ? OFFSET ?
`

var GETUSERBLINKEMAIL = "SELECT kyc.ClientID, kyc.Email, kyc.IdentityNo FROM client_tbl_kyc kyc WHERE kyc.Email = %s"

var MATCHINGDATABTIM = `SELECT c.ClientID FROM client c WHERE c.Email = ? AND c.IDNo = ?
`

var UPDATECLIENTIDBTIM = `
	UPDATE client_tbl_kyc SET ClientIDBTIM = ? WHERE KycID = ?
`
