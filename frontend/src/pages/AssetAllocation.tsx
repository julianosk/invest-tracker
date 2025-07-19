import React, { useEffect, useState } from 'react';
import { fetchInvestments } from '../services/api';
import { Investment } from '../types/investment';

const AssetAllocation: React.FC = () => {
    const [investments, setInvestments] = useState<Investment[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const loadInvestments = async () => {
            try {
                const data = await fetchInvestments();
                setInvestments(data);
            } catch (err) {
                setError('Failed to fetch investments');
            } finally {
                setLoading(false);
            }
        };

        loadInvestments();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div>
            <h1>Asset Allocation</h1>
            <ul>
                {investments.map(investment => (
                    <li key={investment.id}>
                        {investment.name} - {investment.amount} ({investment.assetClass})
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default AssetAllocation;