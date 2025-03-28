# ShouldaBeenBuiltIn

## About

ShouldaBeenBuiltIn enhances your Windows experience by fixing Microsoft's poorly implemented features. Instead of settling for clunky, limited utilities that ship with Windows, we've rebuilt essential tools with the functionality they should have had from the beginning.

Our suite improves critical Windows utilities like the color picker and screenshot tools, eliminating frustrations and adding capabilities that professionals and everyday users have been requesting for years.

## Important Note About Windows Security Warning

When downloading and running this application, Windows may display a security warning ("Windows protected your PC"). This happens because the application is not digitally signed with a certificate from a commercial certificate authority.

**This warning does not mean the application is harmful.**

To run the application:
1. Click "More info" in the warning dialog
2. Click "Run anyway"
3. The application will then start normally

We're working on building reputation with Windows SmartScreen to reduce these warnings over time. Thank you for your understanding!

## Features

- **Enhanced Color Picker**: Quickly select any color on screen with precise control and format conversion
- **Advanced Screenshot Tool**: Capture, annotate, and organize screenshots with ease
- *More tools coming soon...*

## Installation

Download the latest release from the [Releases](https://github.com/ReneSteinmaurer/ShouldaBeenBuiltIn/releases) page.

## Development

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.

## License

[MIT](LICENSE)