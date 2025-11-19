# Nimbus - A humble cloud storage solution

*Currently in active development. More coming soon.*

Nimbus is a mix between lightweight and easy-to-use cloud storage solution and git. It allows users to store, retrieve, and manage their files via CLI. Its build upon a custom binary protocol called RainDrop, designed for efficient file transfer and minimal overhead.

This project is primarily intended for educational purposes, to explore network programming, file I/O, and protocol design. Not recommended for production use.

## Setup

Generate TLS certificates in your home directory:

```bash
mkdir -p ~/.nimbus/certs
cd ~/.nimbus/certs
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/CN=localhost"
```

## Building

Build and install both the client and server:

```bash
# Build and install the client
cd client
go build -o nimbus
go install
cd ..

# Build and install the server
cd server
go build -o nimbus-server
go install
cd ..

# Rename binaries to proper names
mv ~/go/bin/client ~/go/bin/nimbus
mv ~/go/bin/server ~/go/bin/nimbus-server
```

Make sure `~/go/bin` is in your PATH. Add this to your `~/.bashrc`:

```bash
export PATH="$PATH:$HOME/go/bin"
```

Then reload your shell: `source ~/.bashrc`

## Usage

**Start the server:**

```bash
nimbus-server
```

**Send a message from the client:**

```bash
nimbus -send "Your message here"
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
