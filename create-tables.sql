DROP TABLE IF EXISTS producten;
CREATE TABLE producten(
    id          INT AUTO_INCREMENT NOT NULL,
    product     VARCHAR(128) NOT NULL,
    gewicht     INT NOT NULL,
    EAN         INT,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS prijs;
CREATE TABLE prijs(
    id          INT AUTO_INCREMENT NOT NULL,
    lidl        DECIMAL(3,2),
    aldi        DECIMAL(3,2),
    deka        DECIMAL(3,2),
    ah          DECIMAL(3,2),
    kruidvat    DECIMAL(3,2),
    PRIMARY KEY (`id`)
);

INSERT INTO producten
    (product, gewicht)
VALUES
    ('spagetti', 500),
    ('tomaten', 1000),
    ('snoep', 250);

