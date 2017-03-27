# A Client for the [AWS Param Store API Service](https://github.com/kotowick/go-aws-param-store-api-service)

A simple Golang binary that provides a clinet interface for the [AWS Param Store API Service](https://github.com/kotowick/go-aws-param-store-api-service).

## Install

Download the *latest* release. Use it directly, or add it to a server image under /usr/local/bin (or a choice of your liking).

## Usage

**Running the Application**

```
./bin/pac-client -l <landscape> -e <environment> -a <application> -v <version>

Optional Args:

-p <prefix> # Only get params starting with this prefix
-o <output file> # Output params to a file instead of setting them in the shell
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/kotowick/go-aws-param-api-client. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.


## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
