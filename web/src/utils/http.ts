import { ElMessage } from "element-plus"
import { getToken } from "./login"
import { Pagination } from "@/model"

// TODO: validate & update token when fetch
export const useFetch = async (url: string, init?: RequestInit | undefined) => {
  const data: Ref<any> = ref(null)
  const err: Ref<any> = ref(null)
  await fetch(url, init)
    .then((res) => res.json())
    .then((json) => (data.value = json))
    .catch((err) => (err.value = err))

  return { data, err }
}

const useHttp = (baseUrl: string) => (token: string) => {
  const get = (url: string) => {
    return useFetch(`${baseUrl}${url}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
  }
  const post = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  const del = (url: string, id: number) => {
    return useFetch(`${baseUrl}${url}/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
  }
  const put = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  return { get, post, del, put }
}

const baseUrl = '/api/v1'

const _http = useHttp(baseUrl)

export const http = _http(getToken())

export const dialogPost = (token: string) => async (
  api: string,
  _data: any,
  visibleRef: any,
  dataSrc: any
) => {
  const { post } = _http(token)
  const { err } = await post(api, _data)
  if (err.value != null) ElMessage.error(err.value)
  else ElMessage.success('添加成功')
  visibleRef.value = false // 关闭弹窗
  await fetchData(token)(api, dataSrc)
  return err
}

export const fetchData = (token: string) => async (api: string, dataSrc: any) => {
  const { get } = _http(token)
  const { data, err } = await get(api)
  if (err.value != null) ElMessage.warning(err.value)
  dataSrc.value = data.value as any
}

export const deleteData = (token: string) => async (api: string, id: number, dataSrc: any) => {
  const { del } = _http(token)
  const { err } = await del(api, id)
  if (err.value != null) ElMessage.error(err.value)
  else {
    ElMessage.success('删除成功')
    fetchData(token)(api, dataSrc)
  }
}

export const getDataArr = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data: Ref<T[]> = ref([]);
  const get = async (): Promise<T[]> => {
    const { data, err } = await http.get(api);
    if (err) {
      console.error(err);
      return [];
    }
    return (data as Ref<T[]>).value;
  };
  return { data, get };
}

export const getData = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data = ref<T>();
  const get = async (): Promise<T> => {
    const { data, err } = await http.get(api);
    if (err) {
      console.error(err);
      return {} as T;
    }
    return (data as Ref<T>).value;
  };
  return { data, get };
}