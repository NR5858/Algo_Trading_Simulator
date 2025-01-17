'use client'
import { useState } from 'react';

interface SymbolProps {
    symbol: string;
    setSymbol: React.Dispatch<React.SetStateAction<string>>;
}

const Symbol = ({ symbol, setSymbol }: SymbolProps) => {
    return (
        <div className="flex items-center">
            <label className="label mr-4">Symbol:</label>
            <input type="text"
             placeholder="AAPL"
             className="input input-bordered w-full max-w-xs"
             value={symbol}
             onChange={e => setSymbol(e.target.value)} />
        </div>
    );
};

interface ActionProps {
    action: string;
    setAction: React.Dispatch<React.SetStateAction<string>>;
}

const Action = ( {action, setAction }: ActionProps) => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Action:</label>
            <select className="select select-bordered w-full max-w-xs"
             value={action}
             onChange={e => setAction(e.target.value)}>
                <option value="">Select Action</option>
                <option value="Buy">Buy</option>
                <option value="Sell">Sell</option>
            </select>
        </div>
    )
};

interface QuantityProps {
    quantity: number;
    setQuantity: React.Dispatch<React.SetStateAction<number>>;
}

const Quantity = ({quantity, setQuantity}: QuantityProps) => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Quantity:</label>
            <input type="text"
             className="input input-bordered w-full max-w-xs"
             value={quantity}
             onChange={e => setQuantity(Number(e.target.value))} />
        </div>
    );
};

interface PriceTypeProps {
    priceType: string;
    setPriceType: React.Dispatch<React.SetStateAction<string>>;
}

const PriceType = ({ priceType, setPriceType }: PriceTypeProps) => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Price Type:</label>
            <select className="select
             select-bordered
             w-full
             max-w-xs"
             value={priceType}
             onChange={e => setPriceType(e.target.value)}>
                <option value="">Select Price Type</option>
                <option value="Market">Market</option>
                <option value="Limit">Limit</option>
            </select>
        </div>
    )
};

interface DurationProps {
    duration: string;
    setDuration: React.Dispatch<React.SetStateAction<string>>;
}

const Duration = ({ duration, setDuration }: DurationProps) => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Duration:</label>
            <select className="select
             select-bordered
             w-full
             max-w-xs"
             value={duration}
             onChange={e => setDuration(e.target.value)}>
                <option value="">Select Duration</option>
                <option value="Day">Day</option>
                <option value="IOC">IOC</option>
            </select>
        </div>
    )
};

const SendOrderButton = ({ handleSubmit }: {handleSubmit: () => void }) => {
    return (
        <div className="flex items-center mt-5">
            <button className="btn btn-wide" onClick={handleSubmit}>Send Order</button>
        </div>
    );
};

export const OrderForm = () => {
    const [symbol, setSymbol] = useState('');
    const [action, setAction] = useState('');
    const [quantity, setQuantity] = useState(0);
    const [priceType, setPriceType] = useState('');
    const [duration, setDuration] = useState('');

    const handleSubmit = async () => {
        const orderData = {
            symbol,
            action,
            quantity,
            priceType,
            duration,
        };

        // Send orderData to the server
        try {
            const response = await fetch('http://localhost:8080/api/order', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(orderData),
            });
            const result = await response.json();
            console.log('Order created:', result);
        } catch (error) {
            console.error('Error creating order:', error);
        }
    }

    return (
        <div className="flex flex-col items-center justify-center relative top-36">
            <fieldset>
                <Symbol symbol={symbol} setSymbol={setSymbol} />
                <Action action={action} setAction={setAction} />
                <Quantity quantity={quantity} setQuantity={setQuantity} />
                <PriceType priceType={priceType} setPriceType={setPriceType} />
                <Duration duration={duration} setDuration={setDuration} />
                <SendOrderButton handleSubmit={handleSubmit} />
            </fieldset>
        </div>
    );
}