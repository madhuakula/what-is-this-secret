package main

const allRules = `
{
    "rules": [
        {
            "service_name": "slack-token",
            "reg_expression": "(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})"
        },
        {
            "service_name": "rsa-private-key",
            "reg_expression": "-----BEGIN RSA PRIVATE KEY-----"
        },
        {
            "service_name": "ssh-dsa-private-key",
            "reg_expression": "-----BEGIN DSA PRIVATE KEY-----"
        },
        {
            "service_name": "ssh-ec-private-key",
            "reg_expression": "-----BEGIN EC PRIVATE KEY-----"
        },
        {
            "service_name": "pgp-private-key-block",
            "reg_expression": "-----BEGIN PGP PRIVATE KEY BLOCK-----"
        },
        {
            "service_name": "amazon-aws-access-key-id",
            "reg_expression": "AKIA[0-9A-Z]{16}"
        },
        {
            "service_name": "amazon-mws-auth-token",
            "reg_expression": "amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
        },
        {
            "service_name": "aws-api-key",
            "reg_expression": "AKIA[0-9A-Z]{16}"
        },
        {
            "service_name": "facebook-access-token",
            "reg_expression": "EAACEdEose0cBA[0-9A-Za-z]+"
        },
        {
            "service_name": "facebook-oauth",
            "reg_expression": "[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*['|\"][0-9a-f]{32}['|\"]"
        },
        {
            "service_name": "github",
            "reg_expression": "[g|G][i|I][t|T][h|H][u|U][b|B].*['|\"][0-9a-zA-Z]{35,40}['|\"]"
        },
        {
            "service_name": "generic-api-key",
            "reg_expression": "[a|A][p|P][i|I][_]?[k|K][e|E][y|Y].*['|\"][0-9a-zA-Z]{32,45}['|\"]"
        },
        {
            "service_name": "generic-secret",
            "reg_expression": "[s|S][e|E][c|C][r|R][e|E][t|T].*['|\"][0-9a-zA-Z]{32,45}['|\"]"
        },
        {
            "service_name": "google-api-key",
            "reg_expression": "AIza[0-9A-Za-z\\-_]{35}"
        },
        {
            "service_name": "google-cloud-platform-api-key",
            "reg_expression": "AIza[0-9A-Za-z\\-_]{35}"
        },
        {
            "service_name": "google-cloud-platform-oauth",
            "reg_expression": "[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com"
        },
        {
            "service_name": "google-drive-api-key",
            "reg_expression": "AIza[0-9A-Za-z\\-_]{35}"
        },
        {
            "service_name": "google-drive-oauth",
            "reg_expression": "[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com"
        },
        {
            "service_name": "google-gcp-service-account",
            "reg_expression": "\"type\": \"service_account\""
        },
        {
            "service_name": "google-gmail-api-key",
            "reg_expression": "AIza[0-9A-Za-z\\-_]{35}"
        },
        {
            "service_name": "google-gmail-oauth",
            "reg_expression": "[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com"
        },
        {
            "service_name": "google-oauth-access-token",
            "reg_expression": "ya29\\.[0-9A-Za-z\\-_]+"
        },
        {
            "service_name": "google-youtube-api-key",
            "reg_expression": "AIza[0-9A-Za-z\\-_]{35}"
        },
        {
            "service_name": "google-youtube-oauth",
            "reg_expression": "[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com"
        },
        {
            "service_name": "heroku-api-key",
            "reg_expression": "[h|H][e|E][r|R][o|O][k|K][u|U].*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}"
        },
        {
            "service_name": "mailchimp-api-key",
            "reg_expression": "[0-9a-f]{32}-us[0-9]{1,2}"
        },
        {
            "service_name": "mailgun-api-key",
            "reg_expression": "key-[0-9a-zA-Z]{32}"
        },
        {
            "service_name": "password-in-url",
            "reg_expression": "[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@.{1,100}[\"'\\s]"
        },
        {
            "service_name": "paypal-braintree-access-token",
            "reg_expression": "access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"
        },
        {
            "service_name": "picatic-api-key",
            "reg_expression": "sk_live_[0-9a-z]{32}"
        },
        {
            "service_name": "slack-webhook",
            "reg_expression": "https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}"
        },
        {
            "service_name": "stripe-api-key",
            "reg_expression": "sk_live_[0-9a-zA-Z]{24}"
        },
        {
            "service_name": "stripe-restricted-api-key",
            "reg_expression": "rk_live_[0-9a-zA-Z]{24}"
        },
        {
            "service_name": "square-access-token",
            "reg_expression": "sq0atp-[0-9A-Za-z\\-_]{22}"
        },
        {
            "service_name": "square-oauth-secret",
            "reg_expression": "sq0csp-[0-9A-Za-z\\-_]{43}"
        },
        {
            "service_name": "twilio-api-key",
            "reg_expression": "SK[0-9a-fA-F]{32}"
        },
        {
            "service_name": "twitter-access-token",
            "reg_expression": "[t|T][w|W][i|I][t|T][t|T][e|E][r|R].*[1-9][0-9]+-[0-9a-zA-Z]{40}"
        },
        {
            "service_name": "twitter-oauth",
            "reg_expression": "[t|T][w|W][i|I][t|T][t|T][e|E][r|R].*['|\"][0-9a-zA-Z]{35,44}['|\"]"
        }
    ]
}
`
