import type { UserSettings } from "~/routes/auth/usersettings/usersetting";
import { Form, useAsyncValue } from "react-router";
import { Skeleton } from "../Skeleton";

export function SkeletonProfile() {
  return (
    <div className="flex flex-col md:flex-row bg-gray-800 p-6 rounded-lg shadow-lg gap-8">
      <div className="flex flex-col items-center">
        <Skeleton className="w-24 h-24 rounded-full" />
        <Skeleton className="mt-4 h-4 w-20" />
      </div>
      <div className="grid grid-cols-1 md:col-span-2 gap-4 flex-1">
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-20 w-full" />
        <div className="flex justify-between">
          <Skeleton className="h-8 w-32" />
          <Skeleton className="h-8 w-24" />
        </div>
      </div>
    </div>
  );
}

function ProfileForm({ user }: { user: UserSettings }) {
  return (
    <div className="flex flex-col md:flex-row bg-gray-800 p-6 rounded-lg shadow-lg gap-8">
      <div className="flex flex-col items-center">
        <div className="w-24 h-24 rounded-full bg-gray-700 flex items-center justify-center text-2xl font-bold text-white">
          {user.name?.[0]?.toUpperCase()}
        </div>
        <button className="mt-4 text-sm text-red-400 hover:text-red-500">
          Remove photo
        </button>
      </div>

      <Form
        className="grid grid-cols-1 md:col-span-2 gap-4"
        method="post"
        noValidate
      >
        <div className="md:col-span-2">
          <label className="text-sm text-gray-400">Email</label>
          <input
            readOnly
            name="email"
            defaultValue={user.email}
            className="w-full p-2 bg-gray-700 cursor-default rounded-md text-white brightness-50"
          />
          <p className="text-xs text-gray-500 mt-1">
            * Você não pode mudar seu email diretamente
          </p>
        </div>
        <div className="md:col-span-2">
          <label className="text-sm text-gray-400">Contato</label>
          <input
            name="contact"
            defaultValue={user.contact}
            className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
          />
        </div>

        {[
          { label: "Primeiro nome", name: "name", value: user.name },
          { label: "Sobrenome", name: "surname", value: user.surname },
        ].map(({ label, name, value }, i) => (
          <div key={i}>
            <label className="text-sm text-gray-400">{label}</label>
            <input
              name={name}
              defaultValue={value}
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
            120 caracteres faltando
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
      </Form>
    </div>
  );
}

export default function Profile({ user }: { user: UserSettings | null }) {
  if (!user) {
    return <SkeletonProfile />;
  }

  return <ProfileForm user={user} />;
}
