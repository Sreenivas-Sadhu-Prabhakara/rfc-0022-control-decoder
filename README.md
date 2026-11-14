# RFC 22 — Binary control message decoder

**Covered RFC 22. Next RFC 23 — Transmission of Multiple Control Messages.**
The series continues from the [published RFC 21 lab](https://github.com/Sreenivas-Sadhu-Prabhakara/rfc-0021-meeting-lab/commit/9d62ccfc37baeffc8988155109bc65c5acd04a3e).
This lab reuses its Go CLI layout, MIT license and CI workflow.

Vint Cerf's 17 October 1969 memo describes nine binary control commands
for early ARPANET host-to-host links, covering primary and auxiliary
connection setup, acceptance or refusal, pausing, resuming and closing.
An eight-bit opcode determines whether one or two link bytes follow.
The memo permits multiple commands per packet and explains that a
higher-priority CLOSING command may arrive before the final data message.

On 14 November 2026, the RFC Editor catalog listed the Legacy stream,
publication and current status Unknown, and empty updates, updated-by,
obsoletes and obsoleted-by fields. The text describes changes to work
in RFC 11. That reference is separate from the formal metadata.

- [Official RFC 22 text](https://www.rfc-editor.org/info/rfc22/)
- [RFC Editor metadata](https://www.rfc-editor.org/rfc/rfc22.json)

## Applications

- **Protocol trace inspection.** Decode packed commands and identify
  which byte stopped parsing. This is the exercise implemented here.
- **Banking gateway contract tests.** Exercise fixed-width command parsing
  with unknown codes, truncated operands and links that look like opcodes.
- **Distributed stream testing.** Use separate data and control traces
  to test the case where a close directive overtakes in-flight data.

## Run

From the repository root with Go 1.22 or later.

```sh
go run ./cmd/control22 < examples/payload.hex
go run ./cmd/control22 -html < examples/payload.hex > docs/index.html
go test -race ./...
go vet ./...
go build ./...
```

The 11-octet fixture emits five commands at offsets 0, 2, 5, 7 and 9.
The CLI reads whitespace-separated, two-digit hex tokens, with a 64 KiB
input text limit. It emits JSON by default and rejects malformed input
with a nonzero exit. An error produces no partial report. Link parameters
are preserved as unsigned byte values, in order.

The decoder follows the command table's opcode widths. Its labels are
plain-language summaries of the table. In particular, opcode 2 is labeled
Accept primary based on the described exchange of link numbers. The text
rendering calls it DK primary. The lab accepts an already extracted
sequence of command bytes as an educational fixture. Historical packet
headers, marking, checksum, padding, link state and command execution
are outside its scope. GitHub Pages displays a saved Go-generated report.

Tests cover all nine opcode layouts, byte offsets, high-bit link values,
malformed and truncated inputs, the CLI size limit, and JSON and HTML
rendering. CI also regenerates and compares the saved report.
