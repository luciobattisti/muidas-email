import pandas as pd

print("#### Show dataset ####################################")
df = pd.read_csv("data/muidas-newsletter-sheet1.csv")
print(df.head(10))
print(f"Total emails = {len(df)}")

print("#### Lowercase email addresses #######################")
df.loc[:, "Email Address"] = df["Email Address"].apply(lambda x: x.lower())

print("#### Add domain column ###############################")
df.loc[:, "domain"] = df["Email Address"].apply(lambda x: x.split("@")[1])
print(df.head(10))

print("#### Aggregate domains ###############################")
aggregated_domains = df.groupby(["domain"]).size()
aggregated_domains_sorted = aggregated_domains.sort_values(ascending=False)
print(aggregated_domains_sorted)

print("#### Filter out domains with < 10 emails #############")
accepted_domains = list(aggregated_domains.loc[lambda x: x > 10].index)
df_with_accepted_domains = df[df["domain"].isin(accepted_domains)]
print(df_with_accepted_domains.head(10))
print(f"Total emails = {len(df_with_accepted_domains)}")

print("#### Dedup emails ###################################")
df_with_unique_emails = df_with_accepted_domains.groupby(["Email Address"]).tail(1)
print(f"Total unique emails = {len(df_with_unique_emails)}")

print("#### Save to csv ####################################")
df_with_unique_emails.to_csv("data/muidas-newsletter-validated-sheet1.csv", index=False)







