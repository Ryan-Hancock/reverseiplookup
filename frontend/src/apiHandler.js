import axios from "axios";

const axiosInstance = axios.create({
  baseURL: process.env.REACT_BASE_SERVER_URL,
});

const ApiHandler = {
  client: axiosInstance,
  async searchByIP(address) {
    try {
      const response = await this.client.get(`/ip/${address}`);
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  },
  async searchByDomain(address) {
    try {
      const response = await this.client.get(`/domain/${address}`);
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  },
};

export default ApiHandler;
