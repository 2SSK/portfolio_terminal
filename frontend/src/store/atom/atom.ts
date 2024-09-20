import { atom } from "recoil";

import image1 from "/images/wallpapers/image1.jpg";

export const inputState = atom({
  key: "InputState",
  default: "whoami",
});

export const bgState = atom({
  key: "BgState",
  default: image1,
});
