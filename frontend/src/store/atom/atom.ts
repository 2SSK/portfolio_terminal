import { atom } from "recoil";

export const inputState = atom({
  key: "InputState",
  default: "",
});

export const focusInputState = atom<(() => void) | null>({
  key: "focusInputState",
  default: null,
});
