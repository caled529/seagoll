{
  buildGoApplication,
  fetchFromGitHub,
}:
buildGoApplication {
  name = "seagoll";
  src = fetchFromGitHub {
    owner = "caled529";
    repo = "seagoll";
    rev = "7c7b0bb6b4c7315ad24e46d17b86d10e32db62c3";
    hash = "sha256-ZLhL7XKXxbxsWVXGPSNgHyd3baMRW31QopPBUumHSic=";
  };
  pwd = ./.;
}
