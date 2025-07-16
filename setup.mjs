import { execSync, spawn } from "node:child_process";
import * as path from "node:path";
import { existsSync } from "node:fs";
import os from "node:os";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

function run(command, cwd = process.cwd()) {
  // added by vivek
  console.log(`\n[INFO] Running command: "${command}" in ${cwd}`);

  try {
    execSync(command, { stdio: "inherit", cwd });
  } catch (err) {
    // exited by vivek
    console.error(`\n[ERROR] Failed to run command: "${command}". Exiting.`);
    process.exit(1); // Exit with an error code
  }
}

function checkIfExists(command) {
  try {
    // added by vivek
    // Using a platform-agnostic command to check for the tool's existence
    const checkCommand = os.platform() === 'win32' ? `where ${command}` : `command -v ${command}`;
    execSync(checkCommand, { stdio: 'ignore' });
    // execSync(`${command}`, { stdio: "ignore" });

    console.log(`[SUCCESS] Found: ${command}`);
    return true;
  } catch (err) {
    console.log(`[INFO] Not found: ${command}`);
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
        // A more robust check for package managers

        if (checkIfExists('apt')) {
          run('sudo apt update && sudo apt install -y golang-go');
        } else if (checkIfExists('dnf')) {
          run('sudo dnf install -y golang');
        } else if (checkIfExists('pacman')) {
          run('sudo pacman -S --noconfirm go');
        } else {
          console.log('[WARNING] Unsupported Linux distribution. Please install Go manually from https://go.dev/dl/');
        }
        break;


      // if (existsSync("/etc/debian_version")) {
      //   run("sudo apt update && sudo apt install -y golang-go");
      // } else if (existsSync('/etc/fedora-release')) {
      //   run("sudo dnf install -y golang");
      // } else if (existsSync('/etc/arch-release')) {
      //   run("sudo pacman -S --noconfirm go");
      // } else {
      //   console.log("Unsupported Linux distro. Please install Go manually.");
      // }
      // break;
    }
    console.log('[SUCCESS] Go installation command executed. Please ensure it is available in your PATH. You may need to restart your terminal.');
  } catch (err) {
    console.error('[ERROR] Failed to install Go automatically. Please install it manually from https://go.dev/dl/');
    process.exit(1);
  }
}

// Security note: Running scripts directly from the web can be risky.
// This is from the official source, but be aware of the practice.
function installGolangciLint(osType) {
  console.log('[INFO] Installing golangci-lint...');
  switch (osType) {
    case "windows":
      run("winget install --id GolangCI.golangci-lint -e");
      break;
    case "macos":
      run("brew install golangci-lint");
      break;
    case "linux":
      run('curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin');
      break;
    default:
      console.error('[ERROR] Could not install golangci-lint automatically. Please refer to its documentation.');
      process.exit(1);
  }
  console.log('[SUCCESS] golangci-lint installation command executed.');
}


// vivek go air installation
function installAir() {
  if (checkIfExists('go')) {
    console.log('[INFO] Installing air...');
    run('go install github.com/air-verse/air@latest');
    console.log('[SUCCESS] air installation command executed.');
  } else {
    console.error('[ERROR] Go is required to install air. Please install Go first.');
    process.exit(1);
  }
}

// function installAir() {
//   if (checkIfExists("go version")) {
//     run("go install github.com/air-verse/air@latest");
//   } else {
//     console.warn("Go is required to install Air");
//   }
// }

// Using spawn for long-running processes like dev servers is a good practice.
function setupFrontend() {
  console.log('\n[INFO] Setting up frontend...');
  const frontendDir = path.join(__dirname, 'frontend');
  if (existsSync(frontendDir)) {
    if (!existsSync(path.join(frontendDir, 'node_modules'))) {
      run('npm install', frontendDir);
    }
    console.log('[INFO] Starting frontend development server...');
    spawn('npm', ['run', 'dev'], { cwd: frontendDir, stdio: 'inherit', shell: true });
  } else {
    console.error(`[ERROR] Frontend directory not found at: ${frontendDir}`);
    process.exit(1);
  }
}
// function setupFrontend() {
//   const frontend = path.join(__dirname, "frontend");
//   if (existsSync(frontend)) {
//     if (!existsSync(path.join(frontend, "node_modules"))) run("npm install", frontend);
//     spawn("npm", ["run", "dev"], { cwd: frontend, stdio: "inherit", shell: true });
//   }
// }



function setupBackend() {
  console.log('\n[INFO] Setting up backend...');
  const backendDir = path.join(__dirname, 'backend');
  if (existsSync(backendDir)) {
    console.log('[INFO] Starting backend development server with air...');
    spawn('air', [], { cwd: backendDir, stdio: 'inherit', shell: true });
  } else {
    console.error(`[ERROR] Backend directory not found at: ${backendDir}`);
    process.exit(1);
  }
}
// function setupBackend() {
//   const backend = path.join(__dirname, "backend");
//   if (existsSync(backend)) {
//     run("air", backend);
//   }
// }



// --- Main Execution ---
console.log('[INFO] Starting development environment setup...');
const osType = getPlatform();

if (!checkIfExists('go')) {
  installGo(osType);
}

if (!checkIfExists('golangci-lint')) {
  installGolangciLint(osType);
}

if (!checkIfExists('air')) {
  installAir();
}

// if (!checkIfExists("go version")) {
//   installGo(osType);
// }
// if (!checkIfExists("golangci-lint --version")) {
//   installGolangciLint(osType);
// }


// if (!checkIfExists("air -v")) {
//   installAir();
// }

// Start both frontend and backend servers.
// They will run as child processes.
setupFrontend();
setupBackend();


console.log('\n[SUCCESS] All development servers have been started.');
