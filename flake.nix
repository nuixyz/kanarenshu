{
  description = "A minimal TUI application to practise Japanese from the terminal.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      perSystem = {pkgs, ...}: {
        packages.default = pkgs.buildGoModule (finalAttrs: {
          pname = "kanarenshu";
          version = "dev";

          vendorHash = "sha256-ES9+l6aDY8Y38yi4ufw2bpBPCW58L2oSlfXzh1TWGRE=";
          src = ./.;

          meta = {
            description = "A minimal TUI application to practise Japanese from the terminal.";
            homepage = "https://github.com/nuixyz/kanarenshu";
            platforms = pkgs.lib.platforms.unix;
            license = pkgs.lib.licenses.mit;
            mainProgram = "kanarenshu";
          };
        });

        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            go
            gopls
            gotools
          ];
        };
      };
    };
}
