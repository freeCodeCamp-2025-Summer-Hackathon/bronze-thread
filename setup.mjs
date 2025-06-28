import { execSync, spawn } from "node:child_process";
import * as path from "node:path";
import { existsSync } from "node:fs";
import os from "node:os";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

function run(command, cwd = process.cwd()) {
  try {
    execSync(command, { stdio: "inherit", cwd });
  } catch (err) {
    console.error(`Failed to run ${command}: ${err}`);
  }
}

function checkIfExists(command) {
  try {
    execSync(`${command}`, { stdio: "ignore" });
    return true;
  } catch (err) {
    return false;
  }
}

function getPlatform() {
  const platform = os.platform();
  if (platform === "win32") return "windows";
  if (platform === "darwin") return "macos";
  return "linux";
}

function installGo(osType) {
  console.log(`Installing Go on ${osType}...`);
  try {
    switch (osType) {
      case "windows":
        run("winget install --id GoLang.Go -e");
        break;
      case "macos":
        run("brew install go");
        break;
      case "linux":
        if (existsSync("/etc/debian_version")) {
          run("sudo apt update && sudo apt install -y golang-go");
        } else if (existsSync('/etc/fedora-release')) {
          run("sudo dnf install -y golang");
        } else if (existsSync('/etc/arch-release')) {
          run("sudo pacman -S --noconfirm go");
        } else {
          console.log("Unsupported Linux distro. Please install Go manually.");
        }
        break;
    }
  } catch (err) {
    console.error("Failed to install Go automatically. Install it manually from https://go.dev/dl/");
  }
}

function installAir() {
  if (checkIfExists("go")) {
    run("go install github.com/air-verse/air@latest");
  } else {
    console.warn("Go is required to install Air");
  }
}

function setupFrontend() {
  const frontend = path.join(__dirname, "frontend");
  if (existsSync(frontend)) {
    if (!existsSync(path.join(frontend, "node_modules"))) run("npm install", frontend);
    spawn("npm", ["run", "dev"], { cwd: frontend, stdio: "inherit", shell: true });
  }
}

function setupBackend() {
  const backend = path.join(__dirname, "backend");
  if (existsSync(backend)) {
    run("air", backend);
  }
}

const osType = getPlatform();

if (!checkIfExists("go version")) {
  installGo(osType);
}
if (!checkIfExists("air -v")) {
  installAir();
}

setupFrontend();
setupBackend();