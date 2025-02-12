{
  inputs = {
    nixpkgs.url = "github:cachix/devenv-nixpkgs/rolling";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv";
    devenv.inputs.nixpkgs.follows = "nixpkgs";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    self,
    nixpkgs,
    devenv,
    systems,
    ...
  } @ inputs: let
    forEachSystem = nixpkgs.lib.genAttrs (import systems);
  in {
    packages = forEachSystem (system: {
      devenv-up = self.devShells.${system}.default.config.procfileScript;
      devenv-test = self.devShells.${system}.default.config.test;
    });

    devShells =
      forEachSystem
      (system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        default = devenv.lib.mkShell {
          inherit inputs pkgs;
          modules = [
            {
              packages = with pkgs; [
                air
                htmx-lsp
                nodePackages.prettier
              ];
              languages = {
                go.enable = true;
                typescript.enable = true;
                javascript = {
                  enable = true;
                  bun = {
                    enable = true;
                    install.enable = true;
                  };
                };
                rust = {
                  enable = true;
                  toolchain.enable = true;
                };
              };
              services = {
                adminer = {
                  enable = true;
                  listen = "127.0.0.1:8180";
                };
                postgres = {
                  enable = true;
                  listen_addresses = "127.0.0.1";
                  initialDatabases = [
                    {
                      name = "postgres";
                      user = "fady";
                      pass = "fady";
                    }
                    {
                      name = "dogs";
                        #schema = ./Chapter1/native-http/mig/dogs.sql;
                    }
                  ];
                };
              };
            }
          ];
        };
      });
  };
}
