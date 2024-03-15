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
  const put = (url: string, id: number, data: any) => {
    return useFetch(`${baseUrl}${url}/${id}`, {
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

export const dialogPost = async (
  api: string,
  _data: any,
  visibleRef: any,
  dataSrc: any
) => {
  const { post } = http
  const { err } = await post(api, _data)
  if (err.value != null) ElMessage({ message: err.value, type: 'error' })
  else ElMessage({ message: '添加成功', type: 'success' })
  visibleRef.value = false // 关闭弹窗
  await fetchData(api, dataSrc)
  return err
}

export const fetchData = async (api: string, dataSrc: any) => {
  const { get } = http
  const { data, err } = await get(api)
  if (err.value != null) ElMessage({ message: err.value, type: 'error' })
  dataSrc.value = data.value as any
}

export const deleteData = async (api: string, id: number, dataSrc: any) => {
  const { del } = http
  const { err } = await del(api, id)
  if (err.value != null) ElMessage({ message: err.value, type: 'warning' })
  else {
    ElMessage({ message: '删除成功', type: 'success' })
    fetchData(api, dataSrc)
  }
}

export const getDataArr = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data: Ref<T[]> = ref([]);
  const get = async (): Promise<T[]> => {
    const { data, err } = await http.get(api);
    console.log(data.value)
    if (err.value != null) ElMessage({ message: `数据拉取失败：${err}`, type: 'error' });
    return (data as Ref<T[]>).value;
  };
  return { data, get };
}

export const getData = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data = ref({} as T);
  const get = async (): Promise<T> => {
    const { data, err } = await http.get(api);
    console.log(data.value)
    if (err.value != null) ElMessage({ message: `数据拉取失败：${err}`, type: 'error' });
    return (data as Ref<T>).value;
  };
  return { data, get };
}