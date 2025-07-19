import axios from 'axios';
import { Investment } from '../types/investment';

const API_URL = 'http://localhost:5000/api/investments';

export const fetchInvestments = async (): Promise<Investment[]> => {
    const response = await axios.get(API_URL);
    return response.data;
};

export const createInvestment = async (investment: Investment): Promise<Investment> => {
    const response = await axios.post(API_URL, investment);
    return response.data;
};