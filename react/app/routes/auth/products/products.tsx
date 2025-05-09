import makeApi from "~/api/axios";
import type { Route } from "./+types/products";
import { Suspense, useState } from "react";
import { ProductCard } from "~/components/ProductCard";
import { Await, Form, useSearchParams } from "react-router";

export interface LoaderData {
  data: Products[];
  status: number;
  statusText: string | null;
}

export interface Products {
  id: number;
  name: string;
  imageUrl: string;
  price: string;
  discountPercentage: number | null;
  description: string;
  userId: number;
}

async function getProducts(
  request: Request,
  filterText: string | null
): Promise<LoaderData> {
  const api = makeApi(request.headers);
  const params = filterText ? { filterText: filterText } : {};
  return api.get("/products", { params }).then((response) => {
    return {
      data: response.data as Products[],
      status: response.status,
      statusText: response.statusText,
    } as LoaderData;
  });
}

export async function loader({ request }: Route.LoaderArgs) {
  const url = new URL(request.url);
  const filterText = url.searchParams.get("search");
  const productsPromise = getProducts(request, filterText);
  return {
    productsPromise,
  };
}

export default function Products({ loaderData }: Route.ComponentProps) {
  const [searchParams] = useSearchParams();
  const initialSearch = searchParams.get("search") ?? "";
  const [search, setSearch] = useState(initialSearch);

  return (
    <div className="flex">
      <div className="w-96 min-h-screen rounded-lg bg-gray-800 text-white m-10">
        <div className="p-2">
          <Form method="get">
            <div className="relative w-full">
              <input
                id="search"
                name="search"
                type="text"
                placeholder="Pesquisar..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                className="w-full p-2 pr-10 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
              <span className="absolute inset-y-0 right-2 flex items-center text-gray-400 pointer-events-none">
                ⏎
              </span>
            </div>
          </Form>
        </div>
      </div>
      <div className="min-h-screen text-white p-10">
        <div className="mx-auto">
          <div className="bg-gray-800 p-6 rounded-lg shadow-lg">
            {/* todo: should add a skeleton here */}
            <Suspense fallback={<div>Loading...</div>}>
              <Await resolve={loaderData.productsPromise}>
                {(res) => (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6 gap-6 items-stretch">
                    {res.data.map((product) => (
                      <ProductCard
                        key={product.id}
                        id={product.id}
                        name={product.name}
                        price={product.price}
                        discountPercentage={product.discountPercentage}
                        description={product.description}
                        imageUrl={product.imageUrl}
                        href="#"
                      />
                    ))}
                  </div>
                )}
              </Await>
            </Suspense>
          </div>
        </div>
      </div>
    </div>
  );
}
