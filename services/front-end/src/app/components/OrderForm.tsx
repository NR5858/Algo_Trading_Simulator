const Symbol = () => {
    return (
        <div className="flex items-center">
            <label className="label mr-4">Symbol:</label>
            <input type="text"
             placeholder="AAPL"
             className="input input-bordered w-full max-w-xs" />
        </div>
    );
};

const Action = () => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Action:</label>
            <select className="select select-bordered w-full max-w-xs">
                <option>Buy</option>
                <option>Sell</option>
            </select>
        </div>
    )
};

const Quantity = () => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Quantity:</label>
            <input type="text"
             className="input input-bordered w-full max-w-xs" />
        </div>
    );
};

const PriceType = () => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Price Type:</label>
            <select className="select select-bordered w-full max-w-xs">
                <option>Market</option>
                <option>Limit</option>
            </select>
        </div>
    )
};

const Duration = () => {
    return (
        <div className="flex items-center mt-5">
            <label className="label mr-4">Duration:</label>
            <select className="select select-bordered w-full max-w-xs">
                <option>Day</option>
                <option>IOC</option>
            </select>
        </div>
    )
};

export const OrderForm = () => {
    return (
        <div className="flex flex-col items-center justify-center relative top-36">
            <fieldset>
                <Symbol />
                <Action />
                <Quantity />
                <PriceType />
                <Duration />
            </fieldset>
        </div>
    );
}