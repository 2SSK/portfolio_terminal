import { atom } from "recoil";

export const inputState = atom({
  key: "InputState",
  default: { command: "", id: 0 },
});

export const focusInputState = atom<(() => void) | null>({
  key: "focusInputState",
  default: null,
});
