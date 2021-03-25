<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]


<br />
<p align="center">
  <h3 align="center">StructDump</h3>

  <p align="center">
    Simple go structure dumper
    <br />
    <a href="https://pkg.go.dev/github.com/encero/structdump"><strong>Explore the docs »</strong></a>
  </p>
</p>

<!-- ABOUT THE PROJECT -->
## About The Project

StructDump takes go struct and recursively dumps all fields to standard output in "path" format.

The output format is suitable for pasting to excel sheets, documentation, or generally to any text files.

<!-- GETTING STARTED -->
## Installation

```sh
  go get github.com/encero/structdump
```

<!-- USAGE EXAMPLES -->
## Usage

  ```go

  import "github.com/encero/structdump"

  func main() {
    type SimpleType struct {
      AnInt         int
      AnString      string
      AnStringSlice []string
    }

    structdump.Dump(reflect.TypeOf(SimpleType{}))

    // Output:
    // SimpleType.AnInt int
    // SimpleType.AnString string
    // SimpleType.AnStringSlice[] string
  }
  ```

_For more details, please refer to the [Documentation][documentation]_


<!-- CONTRIBUTING -->
## Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/encero/structdump.svg?style=for-the-badge
[contributors-url]: https://github.com/encero/structdump/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/encero/structdump.svg?style=for-the-badge
[forks-url]: https://github.com/encero/structdump/network/members
[stars-shield]: https://img.shields.io/github/stars/encero/structdump.svg?style=for-the-badge
[stars-url]: https://github.com/encero/structdump/stargazers
[issues-shield]: https://img.shields.io/github/issues/encero/structdump.svg?style=for-the-badge
[issues-url]: https://github.com/encero/structdump/issues
[license-shield]: https://img.shields.io/github/license/encero/structdump.svg?style=for-the-badge
[license-url]: https://github.com/encero/structdump/blob/master/LICENSE
[documentation]: https://pkg.go.dev/github.com/encero/structdump