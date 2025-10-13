import { useEffect } from "react";
import { Helmet } from "react-helmet-async";

interface SEOProps {
  title: string;
  description: string;
  keywords?: string;
  canonical?: string;
  noindex?: boolean;
  ogTitle?: string;
  ogDescription?: string;
  ogImage?: string;
  ogType?: string;
  twitterCard?: "summary" | "summary_large_image" | "app" | "player";
  twitterSite?: string;
  twitterCreator?: string;
  schema?: object;
}

const DEFAULT_TITLE = "PayGateway - Gateway de Pagamento Fictício";
const DEFAULT_DESCRIPTION =
  "Sistema completo de gateway de pagamento desenvolvido com arquitetura moderna usando Rust, Go e React. Processamento seguro de transações financeiras.";
const DEFAULT_KEYWORDS =
  "gateway de pagamento, processamento de transações, fintech, rust, go, react, sistema financeiro, api de pagamentos";
const SITE_NAME = "PayGateway";
const SITE_URL = "https://paygateway.exemplo.com";

export const useSEO = (props: SEOProps) => {
  const {
    title,
    description,
    keywords = DEFAULT_KEYWORDS,
    canonical,
    noindex = false,
    ogTitle = title,
    ogDescription = description,
    ogImage = `${SITE_URL}/og-image.png`,
    ogType = "website",
    twitterCard = "summary_large_image",
    twitterSite = "@paygateway",
    twitterCreator = "@eduardomg12",
    schema,
  } = props;

  const fullTitle = title === DEFAULT_TITLE ? title : `${title} | ${SITE_NAME}`;
  const currentUrl = canonical || `${SITE_URL}${window.location.pathname}`;

  const defaultSchema = {
    "@context": "https://schema.org",
    "@type": "WebApplication",
    name: SITE_NAME,
    description: DEFAULT_DESCRIPTION,
    url: SITE_URL,
    applicationCategory: "FinanceApplication",
    operatingSystem: "Web",
    creator: {
      "@type": "Person",
      name: "Eduardo MG",
      url: "https://github.com/EduardoMG12",
    },
    offers: {
      "@type": "Offer",
      price: "0",
      priceCurrency: "BRL",
    },
  };

  const finalSchema = schema || defaultSchema;

  useEffect(() => {
    if (typeof window !== "undefined" && window.gtag) {
      window.gtag("config", "GA_TRACKING_ID", {
        page_title: fullTitle,
        page_location: currentUrl,
      });
    }
  }, [fullTitle, currentUrl]);

  return (
    <Helmet>
      <title>{fullTitle}</title>

      <meta name="description" content={description} />
      <meta name="keywords" content={keywords} />
      <link rel="canonical" href={currentUrl} />

      {noindex && <meta name="robots" content="noindex,nofollow" />}

      <meta property="og:title" content={ogTitle} />
      <meta property="og:description" content={ogDescription} />
      <meta property="og:image" content={ogImage} />
      <meta property="og:type" content={ogType} />
      <meta property="og:url" content={currentUrl} />
      <meta property="og:site_name" content={SITE_NAME} />
      <meta property="og:locale" content="pt_BR" />

      <meta name="twitter:card" content={twitterCard} />
      <meta name="twitter:site" content={twitterSite} />
      <meta name="twitter:creator" content={twitterCreator} />
      <meta name="twitter:title" content={ogTitle} />
      <meta name="twitter:description" content={ogDescription} />
      <meta name="twitter:image" content={ogImage} />

      <script type="application/ld+json">{JSON.stringify(finalSchema)}</script>

      <meta name="theme-color" content="#1a1a1a" />
      <meta name="apple-mobile-web-app-capable" content="yes" />
      <meta name="apple-mobile-web-app-status-bar-style" content="default" />
      <meta name="apple-mobile-web-app-title" content={SITE_NAME} />

      <link rel="preconnect" href="https://fonts.googleapis.com" />
      <link
        rel="preconnect"
        href="https://fonts.gstatic.com"
        crossOrigin="anonymous"
      />
    </Helmet>
  );
};

export const useBasicSEO = (
  title: string,
  description: string,
  keywords?: string,
) => {
  return useSEO({
    title,
    description,
    keywords,
  });
};
