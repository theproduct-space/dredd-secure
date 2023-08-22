/* eslint-disable @typescript-eslint/no-unused-vars */
import {
  useQuery,
  type UseQueryOptions,
  useInfiniteQuery,
  type UseInfiniteQueryOptions,
} from "@tanstack/react-query";
import { useClient } from "../useClient";

export default function useDreddsecureEscrow() {
  const client = useClient();
  const QueryParams = (options: any) => {
    const key = { type: "QueryParams" };
    return useQuery(
      [key],
      () => {
        return client.DreddsecureEscrow.query
          .queryParams()
          .then((res) => res.data);
      },
      options,
    );
  };

  const QueryEscrow = (id: string, options: any) => {
    const key = { type: "QueryEscrow", id };
    return useQuery(
      [key],
      () => {
        const { id } = key;
        return client.DreddsecureEscrow.query
          .queryEscrow(id)
          .then((res) => res.data);
      },
      options,
    );
  };

  const QueryEscrowAll = (query: any, options: any, perPage: number) => {
    const key = { type: "QueryEscrowAll", query };
    return useInfiniteQuery(
      [key],
      ({ pageParam = 1 }: { pageParam?: number }) => {
        const { query } = key;

        query["pagination.limit"] = perPage;
        query["pagination.offset"] = (pageParam - 1) * perPage;
        query["pagination.count_total"] = true;
        return client.DreddsecureEscrow.query
          .queryEscrowAll(query ?? undefined)
          .then((res) => ({ ...res.data, pageParam }));
      },
      {
        ...options,
        getNextPageParam: (lastPage, allPages) => {
          if (
            (lastPage.pagination?.total ?? 0) >
            (lastPage.pageParam ?? 0) * perPage
          ) {
            return lastPage.pageParam + 1;
          } else {
            return undefined;
          }
        },
        getPreviousPageParam: (firstPage, allPages) => {
          if (firstPage.pageParam == 1) {
            return undefined;
          } else {
            return firstPage.pageParam - 1;
          }
        },
      },
    );
  };

  const QueryEscrowsByAddress = (
    address: string,
    query: any,
    options: any,
    perPage: number,
  ) => {
    const key = { type: "QueryEscrowsByAddress", address, query };
    return useInfiniteQuery(
      [key],
      ({ pageParam = 1 }: { pageParam?: number }) => {
        const { address, query } = key;

        query["pagination.limit"] = perPage;
        query["pagination.offset"] = (pageParam - 1) * perPage;
        query["pagination.count_total"] = true;
        return client.DreddsecureEscrow.query
          .queryEscrowsByAddress(address, query ?? undefined)
          .then((res) => ({ ...res.data, pageParam }));
      },
      {
        ...options,
        getNextPageParam: (lastPage, allPages) => {
          if (
            (lastPage.pagination?.total ?? 0) >
            (lastPage.pageParam ?? 0) * perPage
          ) {
            return lastPage.pageParam + 1;
          } else {
            return undefined;
          }
        },
        getPreviousPageParam: (firstPage, allPages) => {
          if (firstPage.pageParam == 1) {
            return undefined;
          } else {
            return firstPage.pageParam - 1;
          }
        },
      },
    );
  };

  const QueryPendingEscrows = (options: any) => {
    const key = { type: "QueryPendingEscrows" };
    return useQuery(
      [key],
      () => {
        return client.DreddsecureEscrow.query
          .queryPendingEscrows()
          .then((res) => res.data);
      },
      options,
    );
  };

  return {
    QueryParams,
    QueryEscrow,
    QueryEscrowAll,
    QueryEscrowsByAddress,
    QueryPendingEscrows,
  };
}
