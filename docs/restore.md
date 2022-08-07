csv!



```go


	_, err := client.CognitoClient.AdminAddUserToGroup(ctx, &cognitoidentityprovider.AdminAddUserToGroupInput{})
	if err {

	}

	_, err := client.CognitoClient.AdminUpdateUserAttributes(ctx, &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: "",
		UserPoolId:     aws.String("us-east-2_SfzBlXTKl"),
		Username:       aws.String(""),
	})
```