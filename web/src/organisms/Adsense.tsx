import React, { useEffect  } from 'react';

function Adsense() {
    useEffect(() => {

        try {
            ((window as any).adsbygoogle = (window as any).adsbygoogle || []).push({});
        }
        catch (e) {

        }
    },[]);

    return (
        <>
            <ins className="adsbygoogle"
                style={{ display: "block" }}
                data-ad-client="ca-pub-3474981323161697"
                data-ad-slot="3107658221"
                data-ad-format="auto"
                data-full-width-responsive="true">
            </ins>
        </>
    );
};

export default Adsense;