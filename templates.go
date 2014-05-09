package erede

const tpl_request_cc = `<Request version="2">
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
			{{if .TipoParcelado}}
			<Instalments>
				<type>{{.TipoParcelado}}</type><number>{{.NumParcelas}}</number>
			</Instalments>
			{{end}}
			{{if .SkipRisk}}{{else}}
			<Risk>
				<Action service="1">
					<MerchantConfiguration>
						<merchant_location>{{.VMLocation}}</merchant_location>
						<channel>{{.VChannel}}</channel>
					</MerchantConfiguration>
					<CustomerDetails>
						<RiskDetails>
							{{if .VuserId}}<user_id>{{.VuserId}}</user_id>{{end}}
							{{if .Vmail}}<email_address>{{.Vmail}}</email_address>{{end}}
							{{if .VIP}}<ip_address>{{.VIP}}</ip_address>{{end}}
						</RiskDetails>
						<PersonalDetails>
							{{if .PFone1}}<telephone>{{.PFone1}}</telephone>{{end}}
							{{if .PFone2}}<telephone_2>{{.PFone2}}</telephone_2>{{end}}
							{{if .PDateBirth}}<date_of_birth >{{.PDateBirth}}</date_of_birth>{{end}}
							{{if .PIDNumber}}<id_number>{{.PIDNumber}}</id_number>
							<id_type>{{.PIDType}}</id_type>{{end}}
							{{if .PFName}}<first_name>{{.PFName}}</first_name>{{end}}
							{{if .PLName}}
							<surname>{{.PLName}}</surname>
							{{end}}
						</PersonalDetails>
						<PaymentDetails>
							<payment_method>CC</payment_method>
						</PaymentDetails>
						{{if .SkipAddressDetails}}{{else}}
						<AddressDetails>
							<address_line1>{{.Address1}}</address_line1>
							<address_line2>{{.Address2}}</address_line2>
							<city>{{.AddressCity}}</city>
							<state_province>{{.AddressState}}</state_province>
							<country>{{.AddressCountry}}</country>
							<zip_code>{{.AddressZIP}}</zip_code>
						</AddressDetails>
						{{end}}
						{{if .SkipShippingDetails}}{{else}}
						<ShippingDetails>
							<delivery_date>{{.DeliveryDate}}</delivery_date>
							<delivery_method>{{.DeliveryMethod}}</delivery_method>
							<address_line1>{{.DeliveryAddress1}}</address_line1>
							<address_line2>{{.DeliveryAddress2}}</address_line2>
							<city>{{.DeliveryCity}}</city>
							<state_province>{{.DeliveryState}}</state_province>
							<country>{{.DeliveryCountry}}</country>
							<zip_code>{{.DeliveryZIP}}</zip_code>
						</ShippingDetails>
						{{end}}
						{{if .SkipOrderDetails}}{{else}}
						<OrderDetails>
							{{if .SkipBillingDetails}}{{else}}
							<BillingDetails>
								<address_line1>{{.Cobranca1}}</address_line1>
								<address_line2>{{.Cobranca2}}</address_line2>
								<city>{{.CobrancaCity}}</city>
								<state_province>{{.CobrancaState}}</state_province>
								<country>{{.CobrancaCountry}}</country>
								<zip_code>{{.CobrancaZIP}}</zip_code>
							</BillingDetails>
							{{end}}
							{{if .SkipLineItems}}{{else}}
							<LineItems>
								{{with .OrderItems}}
								{{range .}}
									<Item>
										<product_code>{{.ProductCode}}</product_code>
										<product_description>{{.ProductDescription}}</product_description>
										<product_category>{{.ProductCategory}}</product_category>
										<order_quantity>{{.OrderQuantity}}</order_quantity>
										<unit_price>{{.UnitPrice}}</unit_price>
										<product_risk>{{.ProductRisk}}</product_risk>
									</Item>
								{{end}}
								{{end}}
							</LineItems>
							{{end}}
						</OrderDetails>
						{{end}}
					</CustomerDetails>
				</Action>
			</Risk>
			{{end}}
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
</Request>`
