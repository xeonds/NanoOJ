import { useI18n } from "vue-i18n"
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
        Authorization: `${token}`,
      },
    })
  }
  const post = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  const del = (url: string, id: number) => {
    return useFetch(`${baseUrl}${url}/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `${token}`,
      },
    })
  }
  const put = (url: string, id: number, data: any) => {
    return useFetch(`${baseUrl}${url}/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  return { get, post, del, put }
}

const baseUrl = '/api/v1'

const _http = useHttp(baseUrl)
const { t } = useI18n()

export const http = _http(getToken())

// deprecated
export const dialogPost = async (
  api: string,
  _data: any,
  visibleRef: any,
  dataSrc: any
) => {
  const { post } = http
  const { err } = await post(api, _data)
  if (err.value != null) ElMessage({ message: err.value, type: 'error' })
  else ElMessage({ message: t('message.add-success'), type: 'success' })
  visibleRef.value = false // 关闭弹窗
  await fetchDataThenUpdateRef(() => http.get(api), dataSrc)
  return err
}

export const fetchDataThenUpdateRef = async (fetch: Function, dataSrc: any) => {
  dataSrc.value = await fetch()
  ElMessage({ message: t('message.data-pull-success'), type: 'success' })
}

// deprecated
export const deleteData = async (api: string, id: number, dataSrc: any) => {
  const { del } = http
  const { err } = await del(api, id)
  if (err.value != null) ElMessage({ message: err.value, type: 'warning' })
  else {
    ElMessage({ message: t('message.delete-success'), type: 'success' })
    fetchDataThenUpdateRef(() => http.get(api), dataSrc)
  }
}

export const getDataArr = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data: Ref<T[]> = ref([]);
  const get = async (): Promise<T[]> => {
    const { data, err } = await http.get(api);
    console.log(data.value)
    if (err.value != null) ElMessage({ message: `${t('message.data-pull-failed')}：${err}`, type: 'error' });
    return (data as Ref<T[]>).value;
  };
  return { data, get };
}

export const getData = <T>(api: string, _pagination = <Pagination>{ pageNum: 1, pageSize: 25, total: 25 }) => {
  const data = ref({} as T);
  const get = async (): Promise<T> => {
    const { data, err } = await http.get(api);
    console.log(data.value)
    if (err.value != null) ElMessage({ message: `${t('message.data-pull-failed')}：${err}`, type: 'error' });
    return (data as Ref<T>).value;
  };
  return { data, get };
}

export const applyData = async (ref: Ref<any>, getter: Function, defaultValue: any) => {
  ref.value = await getter() ?? defaultValue;
}

// TODO: add i18n for default error message
export const handleHttp = (ret: { data: Ref<any>, err: Ref<any> }, _ok: Function, _err?: Function) => {
  if (ret.err.value != null) _err ? _err(ret.err.value) : ElMessage({ message: `${t('message.data-pull-failed')}：${ret.err.value}`, type: 'error' });
  else _ok(ret.data);
}

export const handleArrRefresh = <T>(src: Ref<T[]>, data: T[]) => {
  src.value = data;
  ElMessage({ message: t('message.refresh-success'), type: 'success' });
}