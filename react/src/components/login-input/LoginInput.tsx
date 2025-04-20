import { LoginInputProps } from "../../types/login/LoginInput";

const LoginInput: React.FC<LoginInputProps> = ({ setEmail, setPassword }) => {
  return (
    <div className="container text-center">
      <div className="row">
        <div className="col">
          <div className="form-floating">
            <input
              type="email"
              className="form-control"
              id="emailinput"
              placeholder="Email"
              onChange={(e) => setEmail(e.target.value)}
            />
            <label htmlFor="emailinput">Email</label>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <div className="form-floating">
            <input
              type="Password"
              className="form-control"
              id="passwordinput"
              placeholder="Senha"
              onChange={(e) => setPassword(e.target.value)}
            />
            <label htmlFor="password">Senha</label>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginInput;
