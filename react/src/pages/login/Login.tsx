import React, { useState } from "react";
import LoginInput from "../../components/login-input/LoginInput";
import api from "../../api/axios";
import { AxiosError } from "axios";
import AlertError from "../../components/alert-error/AlertError";
import "../../style/login/login.scss";

const Login: React.FC = () => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [error, setError] = useState<string | null>("");

  const handleLogin = async () => {
    try {
      await api.post("/auth/login", {
        email,
        password,
      });
    } catch (error) {
      const err = error as AxiosError<{ data: string }>;

      if (err.response) {
        const errorMessage = err.response.data;
        setError(errorMessage.data); // Other general error messages
      } else {
        setError("Um erro interno ocorreu. Tente novamente.");
      }
    }
  };

  return (
    <div className="container-fluid">
      <div className="row">
        <div className="col-12 col-md-6 col-lg-3 bg-secondary-subtle">
          <div className="d-flex justify-content-center align-items-center vh-100">
            <div className="col-12">
              <div className="d-flex justify-content-center align-items-center mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="40"
                  height="40"
                  fill="currentColor"
                  className="bi bi-amd me-2"
                  viewBox="0 0 16 16"
                >
                  <path d="m.334 0 4.358 4.359h7.15v7.15l4.358 4.358V0zM.2 9.72l4.487-4.488v6.281h6.28L6.48 16H.2z" />
                </svg>
                <h4 className="mb-0">ECommerce</h4>
              </div>
              <div className="card p-4">
                <div className="text-center mb-3">
                  <h4>Iniciar Sessão</h4>
                </div>

                {
                  <div className="mb-3">
                    <AlertError
                      message={error}
                      onClose={() => setError(null)}
                    />
                  </div>
                }

                <LoginInput setEmail={setEmail} setPassword={setPassword} />

                <div className="d-grid mt-3">
                  <button
                    type="button"
                    className="btn btn-primary"
                    onClick={handleLogin}
                  >
                    Entrar
                  </button>
                </div>

                <hr className="my-4" />
                <p className="text-center mb-2">
                  É um recrutador e gostaria de testar a aplicação?
                </p>
                <div className="d-grid">
                  <button type="button" className="btn btn-outline-primary">
                    Logar como recrutador
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-9 stylish-gradient">
          <div className="container-fluid"></div>
        </div>
      </div>
    </div>
  );
};

export default Login;
