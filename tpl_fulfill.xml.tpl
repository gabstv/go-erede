<Request version="2">
	<Authentication>
		<AcquirerCode><rdcd_pv>{{.User}}</rdcd_pv></AcquirerCode>
		<password>{{.Password}}</password>
	</Authentication>
	<Transaction>
		<HistoricTxn>
			<reference>{{.GatewayReference}}</reference>
			<authcode>{{.AuthCode}}</authcode>
			<method>fulfill</method>
		</HistoricTxn>
	</Transaction>
</Request>