package cmd

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"

	"github.com/0xdeb7ef/spass-manager/pkg/spass"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt a .spass file",
	Long:  `Decrypt a .spass file into a specific format.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Fatal(err)
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal(err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}

		switch format {
		case "raw":
			err = formatRaw(input, output, password)
		case "chrome":
			err = formatChrome(input, output, password)
		case "csv":
			err = formatCSV(input, output, password)
		default:
			log.Fatalf("invalid format: %s", format)
		}

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringP("input", "i", "", "input .spass file")
	decryptCmd.Flags().StringP("output", "o", "", "output filename")
	decryptCmd.Flags().StringP("format", "f", "chrome", "output format [chrome, csv, raw]")
	decryptCmd.Flags().StringP("password", "p", "", "password used to decrypt the .spass file")

	decryptCmd.MarkFlagFilename("input", "spass")

	decryptCmd.MarkFlagRequired("input")
	decryptCmd.MarkFlagRequired("output")
	decryptCmd.MarkFlagRequired("password")
}

func getFile(input, password string) ([]byte, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}

	data, err := spass.Decrypt(file, password)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func formatRaw(input, output, password string) error {
	data, err := getFile(input, password)
	if err != nil {
		return err
	}

	err = os.WriteFile(output, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func formatCSV(input, output, password string) error {
	data, err := getFile(input, password)
	if err != nil {
		return err
	}

	var spass spass.SPASS
	err = spass.Deserialize(data)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	csv := csv.NewWriter(&out)
	csv.Write([]string{"url", "username", "password", "otp", "note"})

	for _, p := range spass.Passwords {
		csv.Write([]string{
			p.Origin_URL,
			p.Username_Value,
			p.Password_Value,
			p.OTP,
			p.Credential_Memo,
		})
	}

	csv.Flush()

	err = os.WriteFile(output, out.Bytes(), 0600)
	if err != nil {
		return err
	}

	return nil
}

func formatChrome(input, output, password string) error {
	data, err := getFile(input, password)
	if err != nil {
		return err
	}

	var spass spass.SPASS
	err = spass.Deserialize(data)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	csv := csv.NewWriter(&out)
	csv.Write([]string{"name", "url", "username", "password", "note"})

	for _, p := range spass.Passwords {
		csv.Write([]string{
			p.Origin_URL,
			p.Origin_URL,
			p.Username_Value,
			p.Password_Value,
			p.Credential_Memo,
		})
	}

	csv.Flush()

	err = os.WriteFile(output, out.Bytes(), 0600)
	if err != nil {
		return err
	}

	return nil
}
