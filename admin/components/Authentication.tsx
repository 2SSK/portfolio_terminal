"use client";
import React, { useState } from "react";
import axios from "axios";

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

const Authentication = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      const response = await axios.post(`${BACKEND_URL}/admin/`, {
        username,
        password,
      });
      alert(response.data.message);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        alert(error.response?.data?.message || "Login failed");
      } else {
        alert("An unexpected error occurred");
      }
    }
  };

  return (
    <div className="text-white w-full flex justify-center">
      <div className="mt-32 border-2 border-green-600 p-8 rounded-lg">
        <div className="flex flex-col items-center justify-center gap-4">
          <InputBox
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Username"
            type="text"
          />
          <InputBox
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            type="password"
          />
          <button
            onClick={handleLogin}
            className="w-full bg-green-600 text-white font-semibold text-xl py-1 rounded-sm active:scale-95"
          >
            login
          </button>
        </div>
      </div>
    </div>
  );
};

const InputBox = ({
  value,
  onChange,
  placeholder,
  type,
}: {
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  placeholder: string;
  type: string;
}) => {
  return (
    <div className="w-full flex justify-center">
      <span className="w-[5em] text-xl font-semibold">{placeholder}</span>
      <input
        type={type}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        className="rounded-sm bg-opacity-6 outline-none text-black px-2 border-2 focus:border-green-500"
      />
    </div>
  );
};

export default Authentication;
