{
  description = "discordle-htmx project flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    templ = {
      url = "github:a-h/templ";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.gitignore.follows = "gitignore";
      inputs.gomod2nix.follows = "gomod2nix";
    };
  };

  outputs = {
    nixpkgs,
    gomod2nix,
    gitignore,
    ...
  } @ inputs: let
    allSystems = [
      "x86_64-linux"
    ];

    forAllSystems = f:
      nixpkgs.lib.genAttrs allSystems (system:
        f {
          inherit system;
          pkgs = import nixpkgs {inherit system;};
        });
  in {
    packages = forAllSystems (
      {
        system,
        pkgs,
        ...
      }: let
        buildGoApplication = gomod2nix.legacyPackages.${system}.buildGoApplication;
        templ = system: inputs.templ.packages.${system}.templ;
      in rec {
        default = discordle-htmx;
        discordle-htmx = buildGoApplication {
          name = "discordle-htmx";
          src = gitignore.lib.gitignoreSource ./.;
          go = pkgs.go;
          # Must be added due to bug https://github.com/nix-community/gomod2nix/issues/120
          pwd = ./.;
          modules = ./gomod2nix.toml;
          preBuild = ''
            ${templ system}/bin/templ generate
          '';
        };
      }
    );
    # `nix develop` provides a shell containing development tools.
    devShell = forAllSystems ({
      system,
      pkgs,
    }: let
      templ = system: inputs.templ.packages.${system}.templ;
    in
      pkgs.mkShell {
        buildInputs = with pkgs; [
          go_1_22
          gomod2nix.legacyPackages.${system}.gomod2nix
          gopls
          air
          tailwindcss
          (templ system)
        ];
      });
  };
}
