package telegramloginwidget_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/LipsarHQ/go-telegram-login-widget"
)

//nolint:funlen // Ignore in test file.
func TestAuthorizationData_Check(t *testing.T) {
	t.Parallel()

	type TestCaseInput struct {
		AuthorizationData *telegramloginwidget.AuthorizationData
		_name             string
		Token             string
	}

	type TestCaseOutput struct {
		Error error
	}

	type TestCase struct {
		Output *TestCaseOutput
		Input  *TestCaseInput
	}

	testCases := []*TestCase{
		{
			Input: &TestCaseInput{
				_name: "ok",
				AuthorizationData: &telegramloginwidget.AuthorizationData{
					AuthDate:  976255200,
					FirstName: "Klim",
					Hash:      "b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95",
					ID:        1,
					LastName:  "Sidorov",
					PhotoURL:  "https://t.me/klimsidorov",
					Username:  "klimsidorov",
				},
				Token: "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX",
			},
			Output: &TestCaseOutput{
				Error: nil,
			},
		},
		{
			Input: &TestCaseInput{
				_name: "error-hash-invalid",
				AuthorizationData: &telegramloginwidget.AuthorizationData{
					AuthDate:  466251600,
					FirstName: "Pavel",
					Hash:      "b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95",
					ID:        1,
					LastName:  "Durov",
					PhotoURL:  "https://t.me/durov",
					Username:  "durov",
				},
				Token: "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX",
			},
			Output: &TestCaseOutput{
				Error: telegramloginwidget.ErrHashInvalid,
			},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.Input._name, func(t *testing.T) {
			t.Parallel()

			require.ErrorIs(t, testCase.Input.AuthorizationData.Check(testCase.Input.Token), testCase.Output.Error)
		})
	}
}

//nolint:funlen // Ignore in test file.
func TestNewFromQuery(t *testing.T) {
	t.Parallel()

	type TestCaseInput struct {
		Query url.Values
		_name string
	}

	type TestCaseOutput struct {
		AuthorizationData *telegramloginwidget.AuthorizationData
		Error             error
	}

	type TestCase struct {
		Output *TestCaseOutput
		Input  *TestCaseInput
	}

	testCases := []*TestCase{
		{
			Input: &TestCaseInput{
				Query: url.Values{
					"auth_date":  []string{"976255200"},
					"first_name": []string{"Klim"},
					"hash":       []string{"b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95"},
					"id":         []string{"1"},
					"last_name":  []string{"Sidorov"},
					"photo_url":  []string{"https://t.me/klimsidorov"},
					"username":   []string{"klimsidorov"},
				},
				_name: "ok",
			},
			Output: &TestCaseOutput{
				AuthorizationData: &telegramloginwidget.AuthorizationData{
					AuthDate:  976255200,
					FirstName: "Klim",
					Hash:      "b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95",
					ID:        1,
					LastName:  "Sidorov",
					PhotoURL:  "https://t.me/klimsidorov",
					Username:  "klimsidorov",
				},
				Error: nil,
			},
		},
		{
			Input: &TestCaseInput{
				Query: url.Values{
					"auth_date":  []string{"466251600"},
					"first_name": []string{"Pavel"},
					"id":         []string{"1"},
					"last_name":  []string{"Durov"},
					"photo_url":  []string{"https://t.me/durov"},
					"username":   []string{"durov"},
				},
				_name: "error-hash-invalid",
			},
			Output: &TestCaseOutput{
				AuthorizationData: nil,
				Error:             telegramloginwidget.ErrHashInvalid,
			},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.Input._name, func(t *testing.T) {
			t.Parallel()

			modelAuthorizationData, err := telegramloginwidget.NewFromQuery(testCase.Input.Query)
			require.ErrorIs(t, err, testCase.Output.Error)
			require.Equal(t, testCase.Output.AuthorizationData, modelAuthorizationData)
		})
	}
}

//nolint:funlen // Ignore in test file.
func TestNewFromURI(t *testing.T) {
	t.Parallel()

	const u = "https://example.com/?"

	type TestCaseInput struct {
		URI   string
		_name string
	}

	type TestCaseOutput struct {
		AuthorizationData *telegramloginwidget.AuthorizationData
		Error             error
	}

	type TestCase struct {
		Output *TestCaseOutput
		Input  *TestCaseInput
	}

	testCases := []*TestCase{
		{
			Input: &TestCaseInput{
				URI: u + url.Values{
					"auth_date":  []string{"976255200"},
					"first_name": []string{"Klim"},
					"hash":       []string{"b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95"},
					"id":         []string{"1"},
					"last_name":  []string{"Sidorov"},
					"photo_url":  []string{"https://t.me/klimsidorov"},
					"username":   []string{"klimsidorov"},
				}.Encode(),
				_name: "ok",
			},
			Output: &TestCaseOutput{
				AuthorizationData: &telegramloginwidget.AuthorizationData{
					AuthDate:  976255200,
					FirstName: "Klim",
					Hash:      "b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95",
					ID:        1,
					LastName:  "Sidorov",
					PhotoURL:  "https://t.me/klimsidorov",
					Username:  "klimsidorov",
				},
				Error: nil,
			},
		},
		{
			Input: &TestCaseInput{
				URI: u + url.Values{
					"auth_date":  []string{"466251600"},
					"first_name": []string{"Pavel"},
					"id":         []string{"1"},
					"last_name":  []string{"Durov"},
					"photo_url":  []string{"https://t.me/durov"},
					"username":   []string{"durov"},
				}.Encode(),
				_name: "error-hash-invalid",
			},
			Output: &TestCaseOutput{
				AuthorizationData: nil,
				Error:             telegramloginwidget.ErrHashInvalid,
			},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.Input._name, func(t *testing.T) {
			t.Parallel()

			modelAuthorizationData, err := telegramloginwidget.NewFromURI(testCase.Input.URI)
			require.ErrorIs(t, err, testCase.Output.Error)
			require.Equal(t, testCase.Output.AuthorizationData, modelAuthorizationData)
		})
	}
}
