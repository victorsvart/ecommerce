import makeApi, { ApiError } from "~/api/axios";
import type { Route } from "./+types/products";
import { Suspense, useEffect, useState } from "react";
import { ProductCard } from "~/components/ProductCard";
import { Await } from "react-router";

export interface LoaderData {
  data: Products[];
  status: number;
  statusText: string | null;
}

export interface Products {
  id: number;
  name: string;
  imageUrl: string;
  userId: number;
}

export async function loader({ request }: Route.LoaderArgs) {
  const api = makeApi(request.headers);

  const productsPromise = api.get("/products").then((response) => {
    return {
      data: response.data as Products[],
      status: response.status,
      statusText: response.statusText,
    };
  });

  return {
    productsPromise,
  };
}

export default function Products({ loaderData }: Route.ComponentProps) {
  return (
    <div className="min-h-screen bg-gray-900 text-white p-10">
      <div className="mx-auto">
        <div className="flex flex-col md:flex-row bg-gray-800 p-6 rounded-lg shadow-lg gap-8">
          <Suspense fallback={<div>Loading...</div>}>
            <Await resolve={loaderData.productsPromise}>
              {(res) =>
                res.data.map((product) => (
                  <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-1 gap-6 items-stretch">
                    <ProductCard
                      key={product.id}
                      id={product.id}
                      name={product.name}
                      description="placeholder description"
                      imageUrl={product.imageUrl}
                      href="#"
                    />
                  </div>
                ))
              }
            </Await>
          </Suspense>
        </div>
      </div>
    </div>
  );
}
