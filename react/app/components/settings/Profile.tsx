import type { UserSettings } from "~/routes/auth/usersettings/usersetting";

export default function Profile({ user }: { user: UserSettings }) {
  return (
    <div className="flex flex-col md:flex-row bg-gray-800 p-6 rounded-lg shadow-lg gap-8">
      <div className="flex flex-col items-center">
        <div className="w-24 h-24 rounded-full bg-gray-700 flex items-center justify-center text-2xl font-bold">
          {user.name?.[0]?.toUpperCase()}{" "}
        </div>
        <button className="mt-4 text-sm text-red-400 hover:text-red-500">
          Remove photo
        </button>
      </div>

      <form className="flex-1 grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="md:col-span-2">
          <label className="text-sm text-gray-400">Email</label>
          <input
            disabled
            value={user.email}
            className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white brightness-50"
          />
          <p className="text-xs text-gray-500 mt-1">
            * Você não pode mudar seu nome de usuário
          </p>
        </div>
        <div className="md:col-span-2">
          <label className="text-sm text-gray-400">Contato</label>
          <input
            value={user.name}
            className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
          />
        </div>

        {[
          { label: "Primeiro nome", value: user.name },
          { label: "Sobrenome", value: user.surname },
        ].map(({ label, value }, i) => (
          <div key={i}>
            <label className="text-sm text-gray-400">{label}</label>
            <input
              value={value}
              className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
            />
          </div>
        ))}

        <div className="md:col-span-2">
          <label className="text-sm text-gray-400">Bio</label>
          <textarea
            rows={3}
            placeholder="Fale sobre você..."
            className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white resize-none"
          />
          <p className="text-right text-xs text-gray-400 mt-1">
            120 characters remaining
          </p>
        </div>

        <div className="md:col-span-2 flex justify-between items-center">
          <a
            href="/profile/public"
            className="text-sm text-blue-400 hover:underline"
          >
            <p>Exibir perfil público</p>
          </a>
          <button
            type="submit"
            className="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md"
          >
            Salvar
          </button>
        </div>
      </form>
    </div>
  );
}
