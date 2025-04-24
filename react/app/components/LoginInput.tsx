import FloatInput from "./FloatInput";

export const LoginInput = () => {
  return (
    <div className="w-full space-y-4">
      <FloatInput id="email" name="email" type="email" label="Email" required />
      <FloatInput
        id="password"
        name="password"
        type="password"
        label="Senha"
        required
      />
    </div>
  );
};
