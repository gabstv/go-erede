<Request version="2">
	<Authentication>
		<AcquirerCode><rdcd_pv>{{.User}}</rdcd_pv></AcquirerCode>
		<password>{{.Password}}</password>
	</Authentication>
	<Transaction>
		<CardTxn>
			<Card>
				<pan>{{.CC}}</pan>
				<expirydate>{{.ExpiryDate}}</expirydate>
				<Cv2Avs><cv2>{{.CCVC2}}</cv2></Cv2Avs>
				<card_account_type>{{.TipoTransacao}}</card_account_type>
			</Card>
			<method>{{.Method}}</method>
		</CardTxn>
		<TxnDetails>
			<merchantreference>{{.MReference}}</merchantreference>
			<capturemethod>ecomm</capturemethod>
			<amount currency="BRL">{{.Amount}}</amount>
			<ThreeDSecure>
				<verify>yes</verify>
				<merchant_url>{{.ThreeDSecure.MerchantURL}}</merchant_url>
				<purchase_desc>{{.ThreeDSecure.PurchaseDesc}}</purchase_desc>
				<purchase_datetime>{{.ThreeDSecure.PurchaseDatetime}}</purchase_datetime>
				{{if .ThreeDSecure.MobileNumber}}<mobile_number>{{.ThreeDSecure.MobileNumber}}</mobile_number>{{end}}
				<Browser>
					<device_category>0</device_category>
					<accept_headers>*/*</accept_headers>
					<user_agent>IE/6.0</user_agent>
				</Browser>
			</ThreeDSecure>
		</TxnDetails>
	</Transaction>
	{{if .ShowUserAgent}}
	<UserAgent>
		<architecture version="2.6.32-5-686">i386-Linux</architecture>
		<language vendor="Sun Microsystems Inc." version="20.1-b02" vm-name="Java HotSpot(TM) Server VM">Java</language>
		<Libraries>
			<lib version="v2-1-3">XMLDocument</lib>
		</Libraries>
	</UserAgent>
	{{end}}
</Request>