{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
  };

  outputs = {self, ...} @ inputs:
    inputs.utils.lib.eachDefaultSystem (
      system: let
        pkgs = import inputs.nixpkgs {
          inherit system;
          overlays = [inputs.gomod2nix.overlays.default];
        };
      in {
        packages = {
          default = pkgs.buildGoApplication {
            name = "seagoll";
            src = fetchGit {
              url = "github:caled529/seagoll.git";
              rev = "7c7b0bb6b4c7315ad24e46d17b86d10e32db62c3";
            };
            pwd = ./.;
          };
        };
        apps = {
          default = {
            type = "app";
            program = "${inputs.self.packages."${pkgs.system}".default}/bin/seagoll";
          };
        };
        devShells = with pkgs; {
          default = mkShell {
            packages = [
              go
              gopls
              inputs.gomod2nix.packages."${pkgs.system}".default
            ];
          };
        };
      }
    );
}
