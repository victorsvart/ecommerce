import React from "react";

type FloatInputProps = {
  id: string;
  label: string;
  type?: string;
  error?: string;
  name: string;
  required?: boolean;
  value?: string;
};

const FloatInput: React.FC<FloatInputProps> = ({
  id,
  label,
  type = "text",
  error,
  name,
  required,
  value = "",
}) => {
  return (
    <div className="relative">
      <input
        defaultValue={value}
        type={type}
        id={id}
        name={name}
        required={required}
        placeholder=" "
        className={`block px-2.5 pb-2.5 pt-4 w-full text-sm text-gray-900 bg-transparent rounded-lg border ${error ? "border-red-500" : "border-gray-300"
          } appearance-none dark:text-white dark:border-gray-600 focus:outline-none focus:ring-0 focus:border-blue-600 peer`}
      />
      <label
        htmlFor={id}
        className="absolute text-sm text-gray-500 dark:text-gray-400 duration-200 transform -translate-y-4 scale-75 top-2 z-10 origin-[0] bg-white dark:bg-gray-900 px-2 peer-focus:px-2 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:top-1/2 peer-focus:top-2 peer-focus:scale-75 peer-focus:-translate-y-4 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto start-1"
      >
        {label}
      </label>
      {error && <p className="text-red-500 text-xs mt-1 text-left">{error}</p>}
    </div>
  );
};

export default FloatInput;
