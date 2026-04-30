import { useEffect, useState } from "react";
import { getProducts } from "../api/productApi";

export default function useProducts(search, page, limit) {
    const [products, setProducts] = useState([]);
    const [total, setTotal] = useState(0);
    const [loading, setLoading] = useState(false);

    const fetchData = async () => {
        setLoading(true);
        try {
            const res = await getProducts({search, page, limit});
            setProducts(res.data.data.data);
            setTotal(res.data.data.total);
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchData()
    }, [search, page])

    return {products, total, loading, refetch: fetchData};
}