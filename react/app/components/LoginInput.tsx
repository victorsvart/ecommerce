import React from "react";
import FloatInput from "./FloatInput";
import type { FieldErrors, UseFormRegister } from "react-hook-form";
type Input = {
  email: string;
  password: string;
};

interface LoginInputProps {
  register: UseFormRegister<Input>;
  errors: FieldErrors<Input>;
}

export const LoginInput: React.FC<LoginInputProps> = ({ register, errors }) => {
  return (
    <div className="w-full space-y-4">
      <FloatInput
        id="email"
        type="email"
        label="Email"
        register={register("email", { required: "Email é obrigatório" })}
        error={errors.email?.message}
      />
      <FloatInput
        id="password"
        type="password"
        label="Senha"
        register={register("password", { required: "Senha é obrigatória" })}
        error={errors.password?.message}
      />
    </div>
  );
};
