import axios from "axios";

const getAccount = async ([url, token]: string[]) => {
  try {
    const res = await axios.get(url, {
      headers: {
        Authorization: `${token}`,
      },
    });
    return res.data.data;
  } catch (error) {
    return error;
  }
};
export { getAccount };
