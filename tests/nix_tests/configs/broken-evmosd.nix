{ pkgs ? import ../../../nix { } }:
let byted = (pkgs.callPackage ../../../. { });
in
byted.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-byted.patch
  ];
})
