{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nur-packages.url = "github:gawakawa/nur-packages";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      nur-packages,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        nurpkgs = nur-packages.legacyPackages.${system};
      in
      {
        devShells.default = nurpkgs.go;
      }
    );
}
