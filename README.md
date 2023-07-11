# Route53-to-ZoneFile Converter

This is a command-line application built in Go that helps you convert AWS Route53 Resource Record Sets to a zone file.
It uses the AWS SDK for Go to list the record sets from a specific hosted zone in your AWS Route53 service and converts
them to a zone file.

## Prerequisites

- Go 1.13 or higher
- AWS account and credentials configured in your system (You can configure it using AWS CLI)
- You must have your hosted zone ID from AWS Route53

## Installation

Clone the repository:

```bash
git clone https://github.com/saga420/route53-to-zonefile.git
```

Navigate to the project directory:

```bash
cd route53-to-zonefile
```

Build the project:

```bash
go build
```

## Usage

After building the project, you can run the application using:

```bash
./route53-to-zonefile --zone-id=<YOUR_HOSTED_ZONE_ID> [--output=<OUTPUT_FILE_NAME>]
```

## Where:

- <YOUR_HOSTED_ZONE_ID>: Your AWS Route53 Hosted Zone ID.
- <OUTPUT_FILE_NAME>: (Optional) The output file name. If not provided, the output will be written to the standard
  output (your terminal).

## For example:

```bash
./route53-to-zonefile --zone-id=Z0123456789ABCDEFABC --output=myzonefile.txt
```

## Note

The program does not include NS (Name Server) and SOA (Start of Authority) records in the output as these are typically
managed by AWS itself in a Route53 Hosted Zone.

## Contributing

We welcome contributions! Please see CONTRIBUTING.md for details on how to contribute.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

If you encounter any issues, please open an issue in this GitHub repository.